package burndownchart

import "3ds-trello-server/models"

type SetBurnDownInterface interface {
	SetDataBurnDownChart(postAuth *models.Auth) (interface{}, error)
}
