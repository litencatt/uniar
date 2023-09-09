package service

import (
	"context"
	"database/sql"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

type Member struct {
	DB      *sql.DB
	Querier repository.Querier
}

func (x *Member) ListMember(ctx context.Context) ([]entity.Member, error) {
	ms, err := x.Querier.GetAllMembers(ctx, x.DB)
	if err != nil {
		return nil, err
	}

	var member []entity.Member
	for _, m := range ms {
		e := entity.Member{
			Name: m.Name,
		}
		member = append(member, e)
	}
	return member, nil
}

func (x *Member) ListProducerMember(ctx context.Context, producerId int64) ([]entity.ProducerMember, error) {
	ms, err := x.Querier.GetProducerMember(ctx, x.DB, producerId)
	if err != nil {
		return nil, err
	}

	var member []entity.ProducerMember
	for _, m := range ms {
		e := entity.ProducerMember{
			MemberID:    m.MemberID,
			Name:        m.Name,
			BondLevel:   m.BondLevelCurent,
			Discography: m.DiscographyDiscTotal,
		}
		member = append(member, e)
	}
	return member, nil
}

func (x *Member) GetMemberByGroup(ctx context.Context, groupId int64) ([]entity.Member, error) {
	ms, err := x.Querier.GetMembersByGroup(ctx, x.DB, groupId)
	if err != nil {
		return nil, err
	}

	var member []entity.Member
	for _, m := range ms {
		e := entity.Member{
			ID:   m.ID,
			Name: m.Name,
		}
		member = append(member, e)
	}
	return member, nil
}

func (x *Member) UpdateProducerMember(ctx context.Context, pm entity.ProducerMember) error {
	err := x.Querier.UpdateProducerMember(ctx, x.DB, repository.UpdateProducerMemberParams{
		ProducerID:           1,
		MemberID:             pm.MemberID,
		BondLevelCurent:      pm.BondLevel,
		DiscographyDiscTotal: pm.Discography,
	})
	if err != nil {
		return err
	}

	return nil
}
