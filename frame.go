package frame

import (
	"github.com/typeck/frame/conf"
	"github.com/typeck/frame/flag"
	"github.com/typeck/frame/log"
	"github.com/typeck/frame/metrics"
	"github.com/typeck/frame/server"
)

func InitSerial(inits... func()) {
	for _, init := range inits {
		init()
	}
}

func InitAll() {
	InitSerial(
		flag.Init,
		conf.Init,
		log.Init,
		metrics.Init,
		server.Init,
		)
}