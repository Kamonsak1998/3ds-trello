package trello

import "3ds-trello-server/models"

type TrelloInterface interface {
	GetCardFinish(idBoard string, token string) ([]*models.Cards, error)
	GetDateLastActivity(idBoard string, token string) ([]*models.Actions, error)
	GetIDList(idBoard string, token string) ([]*models.Lists, error)
}
