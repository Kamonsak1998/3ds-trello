package history

import "3ds-trello-server/models"

type IHistory interface {
	GetHistories(postData *models.Auth, members []*models.Members) (interface{}, error)
}
