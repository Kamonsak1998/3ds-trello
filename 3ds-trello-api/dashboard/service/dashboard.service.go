package service

import (
	"3ds-trello-server/core"
	"3ds-trello-server/dashboard"
	"3ds-trello-server/models"
)

type DashBoardService struct {
	ctx core.IContext
}

func NewDashBoardService(ctx core.IContext) dashboard.DashBoardInterface {
	return &DashBoardService{
		ctx: ctx,
	}
}

func (s *DashBoardService) SaveBoards(board []*models.Boards) error {
	db, err := s.ctx.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	for _, value := range board {
		db.FirstOrCreate(&value, &value)
		db.Model(&board).Where("board_id = ?", value.BoardId).Update("board_name", value.BoardName)
	}
	return nil
}
