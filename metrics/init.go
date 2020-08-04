package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/util"
	"net/http"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", DefaultConfig.Port), mux)
		if err != nil {
			panic(err)
		}
	}()
}
