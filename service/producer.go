package service

import (
	"context"
	"database/sql"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

const (
	IDProviderGoogle = 1
)

type Producer struct {
	DB      *sql.DB
	Querier repository.Querier
}

func (x *Producer) FindProducer(ctx context.Context, identityId string) (entity.Producer, error) {
	p, err := x.Querier.GetProducer(ctx, x.DB, identityId)
	if err != nil {
		return entity.Producer{}, err
	}

	return entity.Producer{
		ID:          p.ID,
		ProviderID:  p.ProviderID,
		IdentityId:  p.IdentityID,
		IsAdmin:     p.IsAdmin == 1,
	}, nil
}

func (x *Producer) RegistProducer(ctx context.Context, identityId string) error {
	err := x.Querier.RegistProducer(ctx, x.DB, repository.RegistProducerParams{
		ProviderID:  IDProviderGoogle,
		IdentityID:  identityId,
	})
	if err != nil {
		return err
	}

	// Regist porducer members
	p, err := x.Querier.GetProducer(ctx, x.DB, identityId)
	if err != nil {
		return err
	}
	members, err := x.Querier.GetAllMembers(ctx, x.DB)
	if err != nil {
		return err
	}
	for _, m := range members {
		if err := x.Querier.RegistProducerMember(ctx, x.DB, repository.RegistProducerMemberParams{
			ProducerID: p.ID,
			MemberID:   m.ID,
		}); err != nil {
			return err
		}
	}
	return nil
}
