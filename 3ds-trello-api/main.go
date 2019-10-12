package main

import (
	dashBoardHttp "3ds-trello-server/dashboard/http"
	"3ds-trello-server/middlewares"
	sprintHttp "3ds-trello-server/sprint/http"
	leaderHttp "3ds-trello-server/leaderboard/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	historyHttp "3ds-trello-server/history/http"

	"log"
)

func main() {
	e := echo.New()
	e.Use(middlewares.Core)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))
	dashBoardHttp.DashboardRoute(e)
	sprintHttp.SprintRoute(e)
	leaderHttp.LeaderBoardRoute(e)
	historyHttp.HistoryRoute(e)
	log.Fatal(e.Start(":9000"))

}
