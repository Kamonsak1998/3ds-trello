package sprint

import (
	"3ds-trello-server/models"
	"time"
)

type DateInterface interface {
	SetDate(postSprint *models.Sprints) (string, error)
	CheckDate(postBoard *models.Auth) (int, interface{}, error)
}

type CardInterface interface {
	SetScoreSize(postPoint *models.Setting) error
	SetScoreCard(postAuth *models.Auth, card []*models.Cards, cardAction []*models.Actions) (string, error)
	CheckScoreSize() (interface{}, error)
}

type SprintInterface interface {
	CheckSprint(idBoard string, cardTime time.Time) (string, error)
}

type ListInterface interface {
	SaveList(postList *models.Setting) error
}
