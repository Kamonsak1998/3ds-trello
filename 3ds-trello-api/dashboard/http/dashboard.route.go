package http

import "github.com/labstack/echo/v4"

func DashboardRoute(e *echo.Echo) {
	controller := NewSetDashboard()
	e.GET("/getdashboard", controller.GetDashboard)
	e.POST("/setmember", controller.SetMembers)
}
