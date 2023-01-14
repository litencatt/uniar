package service

import (
	"context"
	"database/sql"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

type Photgraph struct {
	DB      *sql.DB
	Querier repository.Querier
}

func (x *Photgraph) ListPhotograph(ctx context.Context) ([]entity.Photograph, error) {
	ps, err := x.Querier.GetPhotographListAll(ctx, x.DB)
	if err != nil {
		return nil, err
	}

	var photograph []entity.Photograph
	for _, p := range ps {
		e := entity.Photograph{
			Name: p.Name,
		}
		photograph = append(photograph, e)
	}
	return photograph, nil
}
