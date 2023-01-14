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
