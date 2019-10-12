package service

import (
	"3ds-trello-server/config"
	"3ds-trello-server/core"
	"3ds-trello-server/models"
	"3ds-trello-server/trello"
	"log"
)

type TrelloService struct {
	ctx core.IContext
}

func NewTrelloService(ctx core.IContext) trello.TrelloInterface {
	return &TrelloService{
		ctx: ctx,
	}
}

func GetBoards(token string) ([]*models.Boards, []error) {
	board, err := models.SetMyBoard(config.API, token)
	if err != nil {
		return nil, err
	}
	boards := make([]*models.Boards, 0)
	for _, data := range board {
		boardData := &models.Boards{BoardId: data.ID, BoardName: data.Name, ImageBackground: data.Prefs.BackgroundImage, ColorBackground: data.Prefs.BackgroundColor}
		boards = append(boards, boardData)
	}
	return boards, nil
}

func GetMembers(idBoard string, token string) ([]*models.Members, []error) {
	member, err := models.SetMember(config.API, idBoard, token)
	if err != nil {
		return nil, err
	}
	members := make([]*models.Members, 0)
	for _, value := range member {
		memberData := &models.Members{
			MemberID:   value.ID,
			MemberName: value.FullName,
			AvatarHash: value.AvatarHash,
		}
		members = append(members, memberData)
	}
	return members, nil
}

func (s *TrelloService) GetCardFinish(idBoard string, token string) ([]*models.Cards, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	card := make([]*models.Cards, 0)
	cardBoard, _ := models.SetCardFinish(config.API, token, idBoard)
	if err != nil {
		return nil, err
	}
	idList := []models.Lists{}
	db.Find(&idList)

	if err != nil {
		return nil, err
	}

	for _, value := range cardBoard {
		a := len(value.IDMembers)
		cardData := &models.Cards{}
		for _, valueList := range idList {
			if valueList.ListID == value.IDList {
				if a != 0 {
					cardData = &models.Cards{CardID: value.ID, BoardID: value.IDBoard, MemberID: value.IDMembers[0], CardName: value.Name, CardTime: *value.DateLastActivity}
					card = append(card, cardData)
				} else {
					log.Printf("CardID : %s not have member", value.ID)
				}
			}
		}
	}

	return card, nil
}

func (s *TrelloService) GetDateLastActivity(idBoard string, token string) ([]*models.Actions, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	nameId := []models.Lists{}
	db.Order("list_priority asc").Find(&nameId, "board_id = ?", idBoard)
	index := len(nameId)

	cardAction, _ := models.SetStartTimeListFinish(config.API, idBoard, token)
	Action := make([]*models.Actions, 0)
	Actions := &models.Actions{}
	for _, value := range cardAction {
		if value.Type == "updateCard" {
			if index-1 >= 0 {
			if value.Data.ListAfter.Name == nameId[0].ListName {
				Actions = &models.Actions{CardID: value.Data.Card.ID, Date: value.Date}
			}
		}else if index-1 >= 1 {
			    if value.Data.ListAfter.Name == nameId[1].ListName {
					if value.Data.ListBefore.Name != nameId[0].ListName || value.Data.ListBefore.Name != nameId[2].ListName{
						Actions = &models.Actions{CardID: value.Data.Card.ID, Date: value.Date}
					}
				}
		    }
		} else if value.Type == "createCard" {
			Actions = &models.Actions{CardID: value.Data.Card.ID, Date: value.Date}
		}
		Action = append(Action, Actions)
		Actions = &models.Actions{}
	}
	return Action, nil
}

func (*TrelloService) GetIDList(idBoard string, token string) ([]*models.Lists, error) {
	list, _ := models.SetList(config.API, idBoard, token)
	var listId = &models.Lists{}
	var listIds = []*models.Lists{}
	for _, value := range list {
		listId = &models.Lists{ListID: value.ID, ListName: value.Name}
		listIds = append(listIds, listId)
	}
	return listIds, nil
}
