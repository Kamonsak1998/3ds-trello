package service

import (
	"3ds-trello-server/core"
	"3ds-trello-server/models"
	"3ds-trello-server/sprint"
)

type ListService struct {
	ctx core.IContext
	*models.ScoreSprints
}

func NewListService(ctx core.IContext) sprint.ListInterface {
	return &ListService{
		ctx: ctx,
	}
}
func (s *ListService) SaveList(postList *models.Setting) error {
	db, err := s.ctx.DB()
	if err != nil {
		return err
	}
	var List models.Lists
	var Lists []models.Lists
	var a []string
	defer db.Close()
	idBoard := postList.SprintDate.BoardId
	for i, value := range postList.SelectList {
		List = models.Lists{
			ListID:       value.ListID,
			ListName:     value.ListName,
			BoardID:      idBoard,
			ListPriority: i + 1,
		}
		a = append(a, List.ListID)
		db.FirstOrCreate(&List, "board_id = ? AND list_id = ?", idBoard, value.ListID)
		db.Model(&List).Where("board_id = ? AND list_id = ?", idBoard, value.ListID).Update(map[string]interface{}{
			"list_name":     value.ListName,
			"board_id":      idBoard,
			"list_id":       value.ListID,
			"list_priority": i + 1,
		})
	}
	db.Not("list_id", a).Find(&Lists, "board_id = ?", idBoard)
	for _, values := range Lists {
		ListData := &models.Lists{
			ListID:       values.ListID,
			BoardID:      values.BoardID,
			ListName:     values.ListName,
			ListPriority: values.ListPriority,
		}
		db.Delete(&ListData, &values)
	}
	return nil
}
