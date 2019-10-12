package http

import "github.com/labstack/echo/v4"

func SprintRoute(e *echo.Echo) {
	controller := NewSetSprint()
	e.POST("/checksetting", controller.CheckSettings)
	e.POST("/settingdata", controller.SettingData)
}
