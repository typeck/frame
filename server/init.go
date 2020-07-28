package server

import (
	"fmt"
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/util"
	_ "go.uber.org/automaxprocs"
)

func init() {
	//init http server
	var httpConfig = &HttpConfig{}
	err := conf.Unmarshal("frame.metrics", httpConfig)
	if err != nil {
		fmt.Printf("can't unmarshal http server config:%v, use default config.\n", err)
	}else {
		defaultHttpConfig = httpConfig
	}
	fmt.Printf("init http server config success:%s\n", util.String(httpConfig))
}
