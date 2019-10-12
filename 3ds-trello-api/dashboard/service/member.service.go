package service

import (
	"3ds-trello-server/core"
	"3ds-trello-server/dashboard"
	"3ds-trello-server/models"
)

type MemberService struct {
	ctx core.IContext
}

func NewMemberService(ctx core.IContext) dashboard.MemberInterface {
	return &MemberService{
		ctx: ctx,
	}
}

func (s *MemberService) SaveMembers(member []*models.Members) error {
	db, err := s.ctx.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	for _, value := range member {
		db.FirstOrCreate(&value, &value)
		db.Model(&member).Where("member_id = ?", value.MemberID).Update(map[string]interface{}{
			"member_name": value.MemberName,
			"avatar_hash": value.AvatarHash,
		})
	}
	return nil
}
