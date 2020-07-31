package metrics

import (
	"fmt"
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/util"
)

func Init() {
	var Config = &Config{}
	err := conf.Unmarshal("frame.metrics", Config)
	if err != nil {
		fmt.Printf("can't unmarshal metrics config:%v, use default config.",err)
	}else {
		DefaultConfig = Config
	}
	fmt.Printf("init metrics success:%s.\n\n", util.String(DefaultConfig))
}
