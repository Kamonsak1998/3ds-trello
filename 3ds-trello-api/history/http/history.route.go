package http

import (
	"github.com/labstack/echo/v4"
)

func HistoryRoute(e *echo.Echo) {
	controller := NewHistoryController()
	e.POST("/gethistory", controller.GetHistory)
}
