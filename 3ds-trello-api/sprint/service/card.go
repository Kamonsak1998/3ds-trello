package service

import (
	"3ds-trello-server/core"
	"3ds-trello-server/models"
	"3ds-trello-server/sprint"
	"strings"
	"time"
)

type CardService struct {
	ctx core.IContext
	*models.ScoreSprints
}

func NewCardService(ctx core.IContext) sprint.CardInterface {
	return &CardService{
		ctx: ctx,
	}
}

func (s *CardService) SetScoreSize(postPoint *models.Setting) error {
	db, err := s.ctx.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	Sizes := models.Sizes{Size: []string{"xxs", "xs", "s", "m", "l", "xl", "xxl", "xxxl"}}
	for i, value := range postPoint.ScoreSize.Points {
		sizePoint := models.ScoreSizes{
			SizeID:    Sizes.Size[i],
			SizePoint: value,
		}
		db.FirstOrCreate(&sizePoint, &sizePoint)
		db.Model(&sizePoint).Where("size_id =?", Sizes.Size[i]).Update(map[string]interface{}{
			"size_point": value,
		})
	}
	return nil
}

func (s *CardService) CheckScoreSize() (interface{}, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return nil, err
	}

	sizes := &[]models.ScoreSizes{}
	db.Find(&sizes)
	scoreSize := map[string]interface{}{
		"sizes":  sizes,
		"status": true,
	}
	return scoreSize, nil
}

func (s *CardService) SetScoreCard(postAuth *models.Auth, card []*models.Cards, cardAction []*models.Actions) (string, error) { //Newfeature setpointsize
	db, err := s.ctx.DB()
	sprintSrv := NewSprintService(s.ctx)
	if err != nil {
		return "fail DB", err
	}
	defer db.Close()

	if err != nil {
		return "fail GetCardFinish", err
	}
	Date := &models.Actions{}
	cardData := &models.Cards{}
	score := &[]models.ScoreSizes{}
	deleteCard := []models.Cards{}
	db.Find(&score)
	var a []string
	cards := []models.Cards{}

	for _, value := range card {
		for _, valuesCard := range cardAction {
			if valuesCard.CardID == value.CardID {
				Date = &models.Actions{Date: valuesCard.Date}
				break
			}
		}
		sizeScore := *score
		point := 0.0
		Date := time.Date(Date.Date.Year(), Date.Date.Month(), Date.Date.Day(), 0, 0, 1, 0, time.UTC)
		sprint, err := sprintSrv.CheckSprint(postAuth.BoardID, Date)
		if err != nil {
			return "fail CheckSprint", err
		}
		s := strings.Replace(value.CardName, "(", "", 1)
		char := strings.Split(s, ")")
		char[0] = strings.ToLower(char[0])
		switch char[0] {
		case sizeScore[0].SizeID:
			point = sizeScore[0].SizePoint
		case sizeScore[1].SizeID:
			point = sizeScore[1].SizePoint
		case sizeScore[2].SizeID:
			point = sizeScore[2].SizePoint
		case sizeScore[3].SizeID:
			point = sizeScore[3].SizePoint
		case sizeScore[4].SizeID:
			point = sizeScore[4].SizePoint
		case sizeScore[5].SizeID:
			point = sizeScore[5].SizePoint
		case sizeScore[6].SizeID:
			point = sizeScore[6].SizePoint
		case sizeScore[7].SizeID:
			point = sizeScore[7].SizePoint
		default:
			point = 0.0
		}
		cardData = &models.Cards{
			CardID:     value.CardID,
			BoardID:    value.BoardID,
			MemberID:   value.MemberID,
			CardName:   value.CardName,
			CardPoint:  point,
			CardTime:   Date,
			SprintName: sprint,
		}
		db.FirstOrCreate(&value, "card_id =?", value.CardID)
		if postAuth.BoardID == value.BoardID {
			db.Model(&cardData).Where("board_id = ? AND card_id = ?", postAuth.BoardID, cardData.CardID).
				Update(map[string]interface{}{
					"board_id":    &cardData.BoardID,
					"member_id":   &cardData.MemberID,
					"card_name":   &cardData.CardName,
					"card_time":   &cardData.CardTime,
					"card_point":  &cardData.CardPoint,
					"sprint_name": &cardData.SprintName,
				})
		} else {
			db.Model(&cardData).Where("board_id = ? AND card_id = ?", value.BoardID, cardData.CardID).
				Update(map[string]interface{}{
					"board_id":    &cardData.BoardID,
					"member_id":   &cardData.MemberID,
					"card_name":   &cardData.CardName,
					"card_time":   &cardData.CardTime,
					"card_point":  &cardData.CardPoint,
					"sprint_name": &cardData.SprintName,
				})
		}
		a = append(a, cardData.CardID)
	}

	if a != nil {
		db.Not("card_id", a).Find(&cards, "board_id = ?", postAuth.BoardID)
		for _, values := range cards {
			cardData := &models.Cards{
				CardID:     values.CardID,
				BoardID:    values.BoardID,
				MemberID:   values.MemberID,
				CardName:   values.CardName,
				CardTime:   values.CardTime,
				CardPoint:  values.CardPoint,
				SprintName: values.SprintName,
			}
			db.Delete(&cardData, &values)
		}
	} else {
		db.Find(&deleteCard)
		db.Delete(&deleteCard, "board_id = ?", postAuth.BoardID)
	}

	return "success", nil
}
