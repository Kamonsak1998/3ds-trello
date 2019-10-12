package service

import (
	burndown "3ds-trello-server/burndownchart"
	"3ds-trello-server/core"
	"3ds-trello-server/models"
	"log"
	"math"
)

type SetBurnDownService struct {
	ctx core.IContext
}

func NewBurnDownService(ctx core.IContext) burndown.SetBurnDownInterface {
	return &SetBurnDownService{
		ctx: ctx,
	}
}
func (s *SetBurnDownService) SetDataBurnDownChart(postAuth *models.Auth) (interface{}, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var ScoreTotalSprint, BurnDownChart, ArrayBurnDownChart = &models.DataScoreTotalSprint{}, &models.BurnDownChart{}, &[]models.BurnDownChart{}
	var Cards models.Cards
	var Sprints models.Sprints
	var DataCard []models.Cards
	var DataSprint []models.Sprints
	rows, err := db.Order(`start_date asc`).Model(&Sprints).Where(`board_id = ?`, postAuth.BoardID).Select(
		`sprint_id, sprint_name, start_date, end_date, sprint_day`).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
		log.Println(err)
	}
	for rows.Next() {
		db.ScanRows(rows, &Sprints)
		DataSprint = append(DataSprint, Sprints)
	}
	sprintDay := float64(Sprints.SprintDay)
	rows, err = db.Order(`card_time asc`).Model(&Cards).Where(`board_id = ?`, postAuth.BoardID).Select(
		`card_id, board_id, card_time, card_point, sprint_name `).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
		log.Println(err)
	}
	for rows.Next() {
		db.ScanRows(rows, &Cards)
		DataCard = append(DataCard, Cards)
	}
	for count, valueSprints := range DataSprint {
		sum := 0.0
		for _, valueCards := range DataCard {
			if valueSprints.SprintName == valueCards.SprintName {
				sum += valueCards.CardPoint
			}
		}
		ScoreTotalSprint.StartDate = append(ScoreTotalSprint.StartDate, valueSprints.StartDate.Format("02 January 2006"))
		ScoreTotalSprint.EndDate = append(ScoreTotalSprint.EndDate, valueSprints.EndDate.Format("02 January 2006"))
		ScoreTotalSprint.Title = append(ScoreTotalSprint.Title, valueSprints.SprintName)
		ScoreTotalSprint.Data = append(ScoreTotalSprint.Data, sum)
		Data, DataScore, valueScoreTotal := &models.Data{}, ScoreTotalSprint.Data[count], ScoreTotalSprint.Data[count]
		Data.DataActualBurn = append(Data.DataActualBurn, math.Round(DataScore*100)/100+0)
		times := valueSprints.StartDate.AddDate(0, 0, -1)
		Data.DatePeriod = append(Data.DatePeriod, times.Format("02 Jan"))
		for count2 := 0; count2 < valueSprints.SprintDay; count2++ {
			Day := int(count2) // convert float to int
			time := valueSprints.StartDate.AddDate(0, 0, Day)
			for _, valueCard := range DataCard {
				if valueCard.CardTime == time {
					DataScore = DataScore - valueCard.CardPoint
				}
			}
			Data.DatePeriod = append(Data.DatePeriod, time.Format("02 Jan"))
			Data.DataActualBurn = append(Data.DataActualBurn, math.Round(DataScore*100)/100+0)
		}
		Day := valueScoreTotal / sprintDay
		valueScoreTotal = valueScoreTotal + Day
		var count3 float64
		for count3 = 0; count3 <= sprintDay; count3++ {
			valueScoreTotal = valueScoreTotal - Day
			Data.DataIDealBurn = append(Data.DataIDealBurn, math.Round(valueScoreTotal*100)/100+0)
		}
		Data.StartDate, Data.EndDate, Data.TitleSprint = ScoreTotalSprint.StartDate[count], ScoreTotalSprint.EndDate[count], ScoreTotalSprint.Title[count]
		BurnDownChart.Data = append(BurnDownChart.Data, *Data)
	}
	*ArrayBurnDownChart = append(*ArrayBurnDownChart, *BurnDownChart)
	return BurnDownChart, nil
}
