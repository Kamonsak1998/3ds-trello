package service

import (
	"3ds-trello-server/core"
	"3ds-trello-server/models"
	"3ds-trello-server/sprint"
	"net/http"
)

type DateService struct {
	ctx core.IContext
}

func NewDateService(ctx core.IContext) sprint.DateInterface {
	return &DateService{
		ctx: ctx,
	}
}

func (s *DateService) SetDate(postSprint *models.Sprints) (string, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return "fail", err
	}
	defer db.Close()

	sprint1 := &models.Sprints{
		BoardId:    postSprint.BoardId,
		SprintName: "Sprint 1",
		StartDate:  postSprint.StartDate,
		EndDate:    postSprint.EndDate,
		SprintDay:  postSprint.SprintDay,
	}

	db.FirstOrCreate(&sprint1, "board_id = ? AND sprint_name = ?", postSprint.BoardId, sprint1.SprintName)
	db.Model(&postSprint).Where("board_id = ? AND sprint_name= ?", postSprint.BoardId, sprint1.SprintName).
		Update(map[string]interface{}{
			"start_date": &postSprint.StartDate,
			"end_date":   &postSprint.EndDate,
			"sprint_day": &postSprint.SprintDay,
		})
	db.Delete(&models.Sprints{}, "board_id = ? AND sprint_name != 'Sprint 1'", postSprint.BoardId)

	return "success", nil
}

func (s *DateService) CheckDate(postBoard *models.Auth) (int, interface{}, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer db.Close()

	sprint := &models.Sprints{}
	db.First(&sprint, "board_id = ?", postBoard.BoardID)
	if sprint.SprintId != nil {
		return http.StatusOK, map[string]interface{}{
			"status":    true,
			"message":   "This board is already set time",
			"idBoard":   sprint.BoardId,
			"startDate": sprint.StartDate,
			"endDate":   sprint.EndDate,
			"sprintDay": sprint.SprintDay,
		}, nil
	} else {
		return http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "This board is not set time",
		}, nil
	}

	return 0, nil, nil
}
