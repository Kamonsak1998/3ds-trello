package service

import (
	"3ds-trello-server/core"
	Inleader "3ds-trello-server/leaderboard"
	"3ds-trello-server/models"
	"math"
	"sort"
	"strings"
)

type LeaderBoardService struct {
	ctx     core.IContext
	leaders models.LeaderBoards
	members models.Members
}

func NewLeaderBoardService(ctx core.IContext) Inleader.LeaderInterface {
	return &LeaderBoardService{
		ctx: ctx,
	}
}

func (s *LeaderBoardService) GetLeaderBoardSrv(postAuth *models.Auth, trelloMember []*models.Members) (interface{}, error) {
	db, err := s.ctx.DB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var card, member, leader = &[]models.Cards{}, &[]*models.Members{}, []models.LeaderBoard{}
	db.Find(&card, "board_id = ?", postAuth.BoardID)
	db.Find(&member)
	for _, valueTrello := range trelloMember {
		for _, valueSet := range *member {
			if valueSet.MemberID == valueTrello.MemberID {
				s.members = models.Members{
					MemberID:   valueTrello.MemberID,
					MemberName: valueTrello.MemberName,
					AvatarHash: valueSet.AvatarHash,
				}
			}
		}
		var point float64
		for _, valueCard := range *card {
			if s.members.MemberID == valueCard.MemberID {
				point = point + valueCard.CardPoint
			}
		}
		a, t := models.LeaderBoard{}, strings.Split(s.members.MemberName, " ")
		if s.members.AvatarHash != "" {
			a = models.LeaderBoard{
				MemberNames:  t[0],
				MemberPoints: math.Round(point*100) / 100,
				AvatarsHash:  "https://trello-avatars.s3.amazonaws.com/" + s.members.AvatarHash + "/50.png",
			}
		} else {
			a = models.LeaderBoard{
				MemberNames:  t[0],
				MemberPoints: math.Round(point*100) / 100,
				AvatarsHash:  "https://img.icons8.com/doodle/48/000000/user.png",
			}
		}
		leader = append(leader, a)
	}
	sort.Slice(leader, func(i, j int) bool { return leader[i].MemberPoints > leader[j].MemberPoints })
	s.leaders = models.LeaderBoards{Leader: leader}
	return s.leaders, nil
}
