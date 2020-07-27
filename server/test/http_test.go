package test

import (
	"github.com/labstack/echo/v4"
	"github.com/typeck/frame/server"
	"net/http"
	"testing"
)


func TestServer(t *testing.T) {
	s := server.NewHttp(&server.HttpConfig{
		Port:        1234,
		Host:        "",
		Debug:       true,
		Pprof: 		 true,
	})
	s.GET("/", hello)
	s.Serve()
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
