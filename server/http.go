package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/typeck/frame/log"
	"github.com/typeck/frame/metrics"
	"github.com/typeck/frame/util"
	"net"
	"time"
)

type HttpServer struct {
	*echo.Echo
	config 	*HttpConfig
}

type HttpConfig struct {
	Port 		int				`toml:"port"`
	Host 		string			`toml:"host"`
	Debug 		bool			`toml:"debug"`
	Pprof		bool			`toml:"pprof"`
}

var DefaultHttpServer *HttpServer

var defaultHttpConfig = &HttpConfig{
	Port:  8080,
	Host:  "localhost",
	Debug: true,
	Pprof: true,
}

func (c *HttpConfig)GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func NewHttp(config *HttpConfig) *HttpServer {
	return &HttpServer{
		Echo:echo.New(),
		config: config,
	}

}

func (s *HttpServer) Serve() {
	var err error
	if s.config.Pprof {
		PprofWrapper(s.Echo)
	}

	s.Echo.Listener, err = net.Listen("tcp", s.config.GetAddress())
	if err != nil {
		log.Panic("new listener error", err)
	}
	s.Echo.Debug = s.Debug
	s.Echo.HideBanner = true
	log.Fatal(s.Echo.Start(""))
}

func MetricHandlerFunc(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := util.Str(c.Request().Header.Get("frame_timestamp")).MustInt64()
		err := f(c)
		if err != nil {
			return err
		}
		if t == 0 {
			return nil
		}
		metrics.Histogram.WithLabelValues().Observe(float64(time.Now().UnixNano() - t) / float64(time.Millisecond))
		return nil
	}
}
