package conf

import (
	"fmt"
	"github.com/typeck/frame/flag"
)

func init() {
	var err error
	watch := flag.GetBool("watch")
	configPath := flag.Get("config")
	DefaultConfig, err = NewFromFile(configPath, watch)
	if err != nil {
		panic(err)
	}
	fmt.Printf("init config success, path:%s, watch:%v\n", configPath, watch)
}