package conf

import (
	"fmt"
	"github.com/typeck/frame/flag"
)

func init() {
	var err error
	watch := flag.GetBool("watch")
	configPath := flag.Get("conf")
	DefaultConfig, err = NewFromFile(configPath, watch)
	if err != nil {
		fmt.Printf("init config error,path:%s, watch:%v\n", configPath, watch)
		return
	}
	fmt.Printf("init config success, path:%s, watch:%v\n", configPath, watch)
}