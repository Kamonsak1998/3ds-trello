package http

import (
	"3ds-trello-server/config"
	"3ds-trello-server/core"
	"3ds-trello-server/models"
	sprint "3ds-trello-server/sprint/service"
	trello "3ds-trello-server/trello/service"
	list "3ds-trello-server/sprint/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

type SprintController struct {
}

func NewSetSprint() *SprintController {
	return &SprintController{}
}

func (s *SprintController) CheckSettings(c echo.Context) error {
	cc := c.(core.IContext)
	cardSrv := sprint.NewCardService(cc)
	dateSrv := sprint.NewDateService(cc)
	trelloSrv := trello.NewTrelloService(cc)

	token := c.Request().Header.Get(config.TokenHeader)
	b, _ := ioutil.ReadAll(c.Request().Body)
	postBoard := &models.Auth{}
	json.Unmarshal(b, postBoard)

	_, date, err := dateSrv.CheckDate(postBoard)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "Internal Server Error",
		})
	}
	scoreSize, err := cardSrv.CheckScoreSize()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "Internal Server Error",
		})
	}
	allList, err := trelloSrv.GetIDList(postBoard.BoardID, token)
	list, err := s.getList(cc, allList, postBoard.BoardID)

	settings := models.SettingRes{
		Date:      date,
		ScoreSize: scoreSize,
		List:      list,
	}

	return c.JSON(http.StatusOK, settings)
}

func (s *SprintController) getList(ctx core.IContext, allList []*models.Lists, idBoard string) (interface{}, error) {
	db, err := ctx.DB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	selectLists := []models.Lists{}
	db.Order("list_priority asc").Find(&selectLists, "board_id = ?", idBoard)
	list := map[string]interface{}{
		"list":       allList,
		"selectList": selectLists,
	}

	return list, nil
}

func (*SprintController) SettingData(c echo.Context) error {
	cc := c.(core.IContext)
	setDateSrv := sprint.NewDateService(cc)
	setScoreSrv := sprint.NewCardService(cc)
	setListSrv := list.NewListService(cc)

	postData := models.Setting{}
	body, err := ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(body, &postData)
	err = setScoreSrv.SetScoreSize(&postData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "Internal Server Error",
		})
	}
	status, err := setDateSrv.SetDate(&postData.SprintDate)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  status,
			"message": "Internal Server Error",
		})
	}
	err = setListSrv.SaveList(&postData)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  status,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
	})
}
