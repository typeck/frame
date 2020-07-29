package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
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

var defaultHttpConfig = &HttpConfig{
	Port:  8080,
	Host:  "localhost",
	Debug: true,
	Pprof: true,
}

func (c *HttpConfig)GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func NewHttpWithConfig(config *HttpConfig) *HttpServer {
	return &HttpServer{
		Echo:echo.New(),
		config: config,
	}

}

func NewHttp() *HttpServer {
	return &HttpServer{
		Echo:echo.New(),
		config: defaultHttpConfig,
	}

}

func (s *HttpServer) Serve() error{
	var err error
	if s.config.Pprof {
		PprofWrapper(s.Echo)
	}
	s.Echo.Listener, err = net.Listen("tcp", s.config.GetAddress())
	if err != nil {
		return err
	}
	s.Echo.Debug = s.Debug
	s.Echo.HideBanner = true
	return s.Echo.Start("")
}

func MetricHandlerFunc(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			t := util.Str(c.Request().Header.Get("frame_timestamp")).MustInt64()
			if t > 0 {
				metrics.Histogram.WithLabelValues().Observe(float64(time.Now().UnixNano()-t) / float64(time.Millisecond))
			}

		}()
		return f(c)
	}
}
