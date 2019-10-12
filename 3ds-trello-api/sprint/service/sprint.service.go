package service

import (
	"3ds-trello-server/core"
	"3ds-trello-server/models"
	ISprint "3ds-trello-server/sprint"
	"strconv"
	"time"
)

type SprintService struct {
	ctx core.IContext
}

func NewSprintService(ctx core.IContext) ISprint.SprintInterface {
	return &SprintService{
		ctx: ctx,
	}
}

func (s *SprintService) CheckSprint(idBoard string, cardTime time.Time) (string, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	sprints := &[]models.Sprints{}
	firstSprint := &models.Sprints{}
	lastSprint := &models.Sprints{}
	db.Find(&sprints, "board_id = ?", idBoard)
	for i, value := range *sprints {
		if value.SprintName == "Sprint 1" {
			*firstSprint = value
		}
		if i == len(*sprints)-1 {
			*lastSprint = value
		}
	}
	s.CreateSprint(idBoard, firstSprint, lastSprint, cardTime)
	db.Find(&sprints, "board_id = ?", idBoard)
	for _, value := range *sprints {
		if inTimeSpan(value.StartDate, value.EndDate, cardTime) {
			return value.SprintName, nil
		}
	}

	return "", nil

}

func (s *SprintService) CreateSprint(idBoard string, firstSprint *models.Sprints, lastSprint *models.Sprints, cardTime time.Time) error {
	db, err := s.ctx.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	if inTimeSpan(firstSprint.StartDate, lastSprint.EndDate, cardTime) {
		return nil
	} else if !inTimeSpan(firstSprint.StartDate, lastSprint.EndDate, cardTime) && !cardTime.Before(firstSprint.StartDate) {
		for {
			db.Find(&lastSprint, "board_id = ?", idBoard)
			intSprintName, _ := strconv.Atoi(lastSprint.SprintName[7:])
			NewSprintName := "Sprint " + strconv.Itoa(intSprintName+1)
			newSprint := &models.Sprints{
				SprintId:   nil,
				BoardId:    idBoard,
				SprintName: NewSprintName,
				StartDate:  lastSprint.StartDate.AddDate(0, 0, lastSprint.SprintDay),
				EndDate:    lastSprint.EndDate.AddDate(0, 0, lastSprint.SprintDay),
				SprintDay:  lastSprint.SprintDay,
			}
			db.Create(&newSprint)
			if inTimeSpan(newSprint.StartDate, newSprint.EndDate, cardTime) {
				break
			}
		}
	}
	return nil
}

func inTimeSpan(start, end time.Time, check time.Time) bool {
	end = end.AddDate(0, 0, 1)
	return check.After(start) && check.Before(end)
}
