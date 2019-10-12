package http

import (
	"3ds-trello-server/config"
	"3ds-trello-server/core"
	dashboard "3ds-trello-server/dashboard/service"
	"3ds-trello-server/models"
	"3ds-trello-server/trello/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

type DashBoardController struct {
}

func NewSetDashboard() *DashBoardController {
	return &DashBoardController{}
}

func (*DashBoardController) GetDashboard(c echo.Context) error {
	cc := c.(core.IContext)
	token := c.Request().Header.Get(config.TokenHeader)
	dashBoardSrv := dashboard.NewDashBoardService(cc)
	board, err := service.GetBoards(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	errs := dashBoardSrv.SaveBoards(board)
	if errs != nil {
		log.Println(errs)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, board)
}

func (*DashBoardController) SetMembers(c echo.Context) error {
	cc := c.(core.IContext)
	b, _ := ioutil.ReadAll(c.Request().Body)
	token := c.Request().Header.Get(config.TokenHeader)
	postAuth := models.Auth{}
	json.Unmarshal(b, &postAuth)
	postData := models.Auth{
		BoardID: postAuth.BoardID,
		Token:   token,
	}
	memberSrv := dashboard.NewMemberService(cc)
	member, err := service.GetMembers(postData.BoardID, postData.Token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	errs := memberSrv.SaveMembers(member)
	if errs != nil {
		log.Println(errs)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
	})
}
