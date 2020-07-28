package test

import (
	"github.com/labstack/echo/v4"
	"github.com/typeck/frame/errors"
	"testing"
)


func TestServer(t *testing.T) {
	s := echo.New()
	s.GET("/", hello)
	s.Start(":8080")
}

// Handler
func hello(c echo.Context) error {
	return errors.New("wrong")
	//return c.String(http.StatusOK, "Hello, World!")
}
