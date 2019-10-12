package service

import (
	"3ds-trello-server/core"
	IHis "3ds-trello-server/history"
	"3ds-trello-server/models"
	"math"
)

// HistoryService ...
type HistoryService struct {
	ctx       core.IContext
	histories models.History
}

// NewHistoryService ...
func NewHistoryService(ctx core.IContext) IHis.IHistory {
	return &HistoryService{
		ctx: ctx,
	}
}

func (s *HistoryService) GetHistories(postData *models.Auth, members []*models.Members) (interface{}, error) {
	s.SetScoreSprints(postData.BoardID, members)
	s.SetScoreTotal(postData.BoardID, members)
	return s.histories, nil
}

func (s *HistoryService) SetScoreSprints(idBoard string, members []*models.Members) error {
	db, err := s.ctx.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	cards := &[]models.Cards{}
	sprints := &[]models.Sprints{}
	db.Find(&sprints, "board_id = ?", idBoard)
	for _, sprintValue := range *sprints {
		scoreOfSprints := &models.ScoreOfSprint{}
		var scoreSprintData []float64
		var scoreSprintNames []string
		for _, memberValue := range members {
			sum := 0.0
			db.Find(&cards, "board_id = ? AND member_id = ? AND sprint_name = ?", idBoard, memberValue.MemberID, sprintValue.SprintName)
			for _, cardValue := range *cards {
				sum += cardValue.CardPoint
			}
			scoreSprintData = append(scoreSprintData, math.Round(sum*100)/100)
			scoreSprintNames = append(scoreSprintNames, memberValue.MemberName)
		}
		scoreOfSprints = &models.ScoreOfSprint{
			Title:     sprintValue.SprintName,
			StartDate: sprintValue.StartDate.Format("02 January 2006"),
			EndDate:   sprintValue.EndDate.Format("02 January 2006"),
			Data:      scoreSprintData,
			Name:      scoreSprintNames,
		}
		s.histories.ScoreOfSprints = append(s.histories.ScoreOfSprints, *scoreOfSprints)
	}

	return nil
}

func (s *HistoryService) SetScoreTotal(idBoard string, members []*models.Members) error {
	db, err := s.ctx.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	cards := &[]models.Cards{}
	sprints := &[]models.Sprints{}
	for _, memberValue := range members {
		sum := 0.0
		db.Find(&sprints, "board_id = ?", idBoard)
		for _, sprintValue := range *sprints {
			db.Find(&cards, "board_id = ? AND member_id = ? AND sprint_name = ?", idBoard, memberValue.MemberID, sprintValue.SprintName)
			for _, cardValue := range *cards {
				sum += cardValue.CardPoint
			}
		}
		s.histories.ScoreTotal.Name = append(s.histories.ScoreTotal.Name, memberValue.MemberName)
		s.histories.ScoreTotal.Data = append(s.histories.ScoreTotal.Data, math.Round(sum*100)/100)
	}

	return nil
}
