package http

import (
	"3ds-trello-server/config"
	"3ds-trello-server/core"
	burnDownSrv "3ds-trello-server/burndownchart/service"
	historySrv "3ds-trello-server/history/service"
	sprintSrv "3ds-trello-server/sprint/service"
	"3ds-trello-server/models"
	trello "3ds-trello-server/trello/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

// HistoryController ...
type HistoryController struct {
}

// NewHistoryController ...
func NewHistoryController() *HistoryController {
	return &HistoryController{}
}

func (*HistoryController) GetHistory(c echo.Context) error {
	cc := c.(core.IContext)
	token := c.Request().Header.Get(config.TokenHeader)
	b, _ := ioutil.ReadAll(c.Request().Body)
	postAuth := models.Auth{}
	json.Unmarshal(b, &postAuth)
	postData := models.Auth{
		BoardID: postAuth.BoardID,
		Token:   token,
	}
	cardSrv := sprintSrv.NewCardService(cc)
	hisSrv := historySrv.NewHistoryService(cc)
	burnSrv := burnDownSrv.NewBurnDownService(cc)
	trelloSrv := trello.NewTrelloService(cc)
	card, _ := trelloSrv.GetCardFinish(postAuth.BoardID, token)
	cardAction, _ := trelloSrv.GetDateLastActivity(postAuth.BoardID, token)
	_, err := cardSrv.SetScoreCard(&postData, card, cardAction)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "Internal Server Error",
		})
	}
	members, errs := trello.GetMembers(postData.BoardID, postData.Token)
	if errs != nil {
		log.Println(errs)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	histories, err := hisSrv.GetHistories(&postData, members)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "Internal Server Error",
		})
	}
	burnDown, err := burnSrv.SetDataBurnDownChart(&postData)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"histories": histories,
		"burnDown":  burnDown,
	})
}
