package log

import (
	"fmt"
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/util"
)

func init() {
	var config = &Config{}
	err := conf.Unmarshal("frame.log", config)
	if err != nil {
		fmt.Printf("can't unmarshal log config:%v, use default config.\n",err)
		config = DefaultConfig()
	}
	DefaultLogger = New(config)
	fmt.Printf("init log success:%s\n", util.String(config))
}
