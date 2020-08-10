package log

import (
	"fmt"
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/util"
	"go.uber.org/zap"
)

func Init() {
	var config = &Config{}
	err := conf.Unmarshal("frame.log", config)
	if err != nil {
		fmt.Printf("can't unmarshal log config:%v, use default config.\n",err)
		config = DefaultConfig()
	}
	DefaultLogger = New(config, zap.AddCaller(), zap.AddCallerSkip(3))
	fmt.Printf("init log success:%s\n\n", util.String(config))
}
