package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pelletier/go-toml"
	"github.com/typeck/frame/errors"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Config struct {
	mux 		sync.RWMutex
	Tree 		*toml.Tree
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
	c.Tree = tree
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
	node := c.Tree.GetPath(strings.Split(key, "."))
	if node == nil {
		return nil
	}
	c.mux.RUnlock()
	if v, ok := node.(*toml.Tree); ok {
		clone := &Config{Tree: v}
		c.cache.Store(key, clone)
		return clone
	}
	c.cache.Store(key, node)
	return node
}

func (c *Config) GetInt(key string) int {
	return int(c.GetInt64(key))
}

func (c *Config) GetInt64(key string) int64 {
	value := c.Get(key)
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case string:
		tmp,_ := strconv.Atoi(v)
		return int64(tmp)
	default:
		return 0
	}
}

func (c *Config) GetStr(key string) string {
	value := c.Get(key)
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	case int64:
		return strconv.Itoa(int(v))
	case int:
		return strconv.Itoa(v)
	default:
		return ""
	}
}

func (c *Config) Unmarshal(key string, v interface{}) error {
	if key == "" {
		return c.Tree.Unmarshal(v)
	}
	node, ok := c.Get(key).(*Config);
	if !ok {
		return errors.New("can't unmarshal.")
	}
	return node.Tree.Unmarshal(v)
}

func (c *Config) String() string {
	result, _ := c.Tree.ToTomlString()
	return result
}