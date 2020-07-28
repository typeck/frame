package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"path/filepath"
	"sync"
)

func (c *Config) Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	eventWg := sync.WaitGroup{}
	eventWg.Add(1)
	configFile := filepath.Clean(c.filePath)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					eventWg.Done()
					return
				}
				const writeOrCreateMask = fsnotify.Write | fsnotify.Create
				if filepath.Clean(event.Name) == configFile && event.Op&writeOrCreateMask != 0 {
					err := c.LoadFromFile()
					if err != nil {
						fmt.Printf("error reading config file: %v\n", err)
					}
					if c.onChange != nil {
						c.onChange(event)
					}
				}else if filepath.Clean(event.Name) == configFile &&
					event.Op&fsnotify.Remove&fsnotify.Remove != 0 {
					eventWg.Done()
					return
				}

			case err, ok := <-watcher.Errors:
				if ok { // 'Errors' channel is not closed
					fmt.Printf("watcher error: %v\n", err)
				}
				eventWg.Done()
				return
			}
		}
	}()
	watcher.Add(c.filePath)
	eventWg.Wait()
}
