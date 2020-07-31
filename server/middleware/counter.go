package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/typeck/frame/metrics"
	"strconv"
	"time"
)

func Counter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t := time.Now()
			c.Request().Header.Set("X-Frame-Timestamp", strconv.Itoa(int(t.UnixNano())))
			metrics.Counter.WithLabelValues(c.Path(), metrics.COUNT, "req").Inc()
			return nil
		}
	}
}
