package server

import (
	"fmt"
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/util"
	_ "go.uber.org/automaxprocs"
)

func Init() {
	//init http server
	var httpConfig = &HttpConfig{}
	err := conf.Unmarshal("frame.http", httpConfig)
	if err != nil {
		fmt.Printf("can't unmarshal http server config:%v, use default config.\n", err)
	}else {
		defaultHttpConfig = httpConfig
	}
	fmt.Printf("init http server config success:%s\n\n", util.String(httpConfig))
}
