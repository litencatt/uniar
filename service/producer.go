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
	return nil
}
