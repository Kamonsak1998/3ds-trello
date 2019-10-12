package dashboard

import "3ds-trello-server/models"

type DashBoardInterface interface {
	SaveBoards(board []*models.Boards) error
}
type MemberInterface interface {
	SaveMembers(board []*models.Members) error
}
