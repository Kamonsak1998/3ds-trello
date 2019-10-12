package http

import "github.com/labstack/echo/v4"

func LeaderBoardRoute(e *echo.Echo) {
	controller := NewLeaderBoardController()
	e.POST("/getleaderboard", controller.GetLeaderBoard)

}
