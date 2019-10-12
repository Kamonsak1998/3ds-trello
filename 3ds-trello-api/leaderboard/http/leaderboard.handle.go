package http

import (
	"3ds-trello-server/config"
	"3ds-trello-server/core"
	leaderSrv "3ds-trello-server/leaderboard/service"
	"3ds-trello-server/models"
	"3ds-trello-server/sprint/service"
	trello "3ds-trello-server/trello/service"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LeaderBoardController struct {
}

func NewLeaderBoardController() *LeaderBoardController {
	return &LeaderBoardController{}
}

func (*LeaderBoardController) GetLeaderBoard(c echo.Context) error {
	cc := c.(core.IContext)
	token := c.Request().Header.Get(config.TokenHeader)
	b, err := ioutil.ReadAll(c.Request().Body)
	postAuth := models.Auth{}
	json.Unmarshal(b, &postAuth)
	postData := models.Auth{
		BoardID: postAuth.BoardID,
		Token:   token,
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Internal Server Error",
			"message": err,
		})
	}
	trelloSrv := trello.NewTrelloService(cc)
	card, err := trelloSrv.GetCardFinish(postData.BoardID, postData.Token)
	cardAction, err := trelloSrv.GetDateLastActivity(postData.BoardID, postData.Token)
	sprintSrv := service.NewCardService(cc)
	_, err = sprintSrv.SetScoreCard(&postData, card, cardAction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Internal Server Error",
			"message": err,
		})
	}
	leaderSrv := leaderSrv.NewLeaderBoardService(cc)
	member, errs := trello.GetMembers(postData.BoardID, postData.Token)
	if errs != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Internal Server Error",
			"message": errs,
		})
	}
	allData, err := leaderSrv.GetLeaderBoardSrv(&postData, member)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Internal Server Error",
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, allData)
}
