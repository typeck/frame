package frame

import (
	"fmt"
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/flag"
	"github.com/typeck/frame/log"
	"github.com/typeck/frame/metrics"
	"github.com/typeck/frame/util"
	_ "go.uber.org/automaxprocs"
)

func init() {
	var err error
	//init config
	watch := flag.GetBool("watch")
	configPath := flag.Get("config")
	conf.DefaultConfig, err = conf.NewFromFile(configPath, watch)
	if err != nil {
		panic(err)
	}
	fmt.Printf("init config success, path:%s, watch:%v", configPath, watch)
	//init log
	var logConfig = &log.Config{}
	err = conf.Unmarshal("log", logConfig)
	if err != nil {
		fmt.Printf("can't unmarshal log config:%v, use default config.",err)
		logConfig = log.DefaultConfig()
	}
	log.DefaultLogger = log.New(logConfig)
	fmt.Printf("init log success:%s", util.String(logConfig))

	//init metric
	var config = &metrics.Config{}
	err = conf.Unmarshal("metrics", config)
	if err != nil {
		fmt.Printf("can't unmarshal metrics config:%v, use default config.",err)
	}else {
		metrics.DefaultConfig = config
	}
	fmt.Printf("init metrics success:%s", util.String(metrics.DefaultConfig))

}
