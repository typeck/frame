package test

import (
	"github.com/labstack/echo/v4"
	"testing"
)


func TestServer(t *testing.T) {
	s := echo.New()
	s.POST("/", hello)
	s.Start(":8080")
}
type User struct {
	Name 	string
}
// Handler
func hello(c echo.Context) error {
	var user = User{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(400, "err")
	}
	return test(c)

}

func test(c echo.Context) error{
	return c.String(200, "98888")
}
