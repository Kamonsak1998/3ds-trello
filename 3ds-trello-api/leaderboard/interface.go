package leaderboard

import "3ds-trello-server/models"

type LeaderInterface interface {
	GetLeaderBoardSrv(postAuth *models.Auth, members []*models.Members) (interface{}, error)
}
