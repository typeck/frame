package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pelletier/go-toml"
	"os"
	"strings"
	"sync"
)

type Config struct {
	mux 		sync.RWMutex
	tree 		*toml.Tree
	cache 		sync.Map
	watch 		bool
	filePath 	string
	onChange 	func(fsnotify.Event)
}

var DefaultConfig *Config

var defaultOnchange = func(e fsnotify.Event) {
	fmt.Printf("config file changed:%s", e.Name)
}

//load config file into config struct.
func NewFromFile(path string, watch bool)(*Config, error) {

	config := &Config{
		watch:    watch,
		filePath: path,
	}
	err := config.LoadFromFile()
	return config, err
}

func (c *Config) LoadFromFile() error{
	file, err := os.Open(c.filePath)
	if err != nil {
		return  err
	}
	defer file.Close()
	tree, err := toml.LoadReader(file)
	if err != nil {
		return  err
	}
	c.mux.Lock()
	c.tree = tree
	c.mux.Unlock()
	c.cache = sync.Map{}
	if c.onChange == nil {
		c.onChange = defaultOnchange
	}
	if c.watch {
		go c.Watch()
	}
	return nil
}

func (c *Config) Get(key string) interface{} {
	if key == "" {
		return c
	}
	if v, ok := c.cache.Load(key); ok {
		return v
	}
	c.mux.RLock()
	node := c.tree.GetPath(strings.Split(key, "."))
	c.mux.RUnlock()
	if v, ok := node.(*toml.Tree); ok {
		clone := &Config{tree: v}
		c.cache.Store(key, clone)
		return clone
	}
	c.cache.Store(key, node)
	return node
}

func (c *Config) GetInt(key string) int {
	return c.Get(key).(int)
}

func (c *Config) GetInt64(key string) int64 {
	return c.Get(key).(int64)
}

func (c *Config) GetStr(key string) string {
	return c.Get(key).(string)
}

func (c *Config) Unmarshal(key string, v interface{}) error {
	if key == "" {
		return c.tree.Unmarshal(v)
	}
	node := c.Get(key).(*Config)
	return node.tree.Unmarshal(v)
}

func (c *Config) String() string {
	result, _ := c.tree.ToTomlString()
	return result
}