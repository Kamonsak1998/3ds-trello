package middlewares

import (
	"3ds-trello-server/core"
	"github.com/labstack/echo/v4"
)

func Core(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &core.Context{c}
		return h(cc)
	}
}
