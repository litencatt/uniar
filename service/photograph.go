package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

type Photgraph struct {
	DB      *sql.DB
	Querier repository.Querier
}

type UpdatePhotographParams struct {
	Name         string
	GroupID      int64
	PhotoType    string
	Abbreviation string
	NameForOrder string
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

func (x *Photgraph) GetPhotographByGroup(ctx context.Context, groupId int64) ([]entity.Photograph, error) {
	ps, err := x.Querier.GetPhotographByGroupId(ctx, x.DB, groupId)
	if err != nil {
		return nil, err
	}

	var photograph []entity.Photograph
	for _, p := range ps {
		e := entity.Photograph{
			ID:   p.ID,
			Name: p.Name,
		}
		photograph = append(photograph, e)
	}
	return photograph, nil
}

func (x *Photgraph) GetSsrPlusReleasedPhotographList(ctx context.Context) ([]entity.Photograph, error) {
	ps, err := x.Querier.GetSsrPlusReleasedPhotographList(ctx, x.DB)
	if err != nil {
		return nil, err
	}

	var photograph []entity.Photograph
	for _, p := range ps {
		e := entity.Photograph{
			ID: p,
		}
		photograph = append(photograph, e)
	}
	return photograph, nil
}

func (x *Photgraph) GetByID(ctx context.Context, id int64) (*entity.Photograph, error) {
	p, err := x.Querier.GetPhotographById(ctx, x.DB, id)
	if err != nil {
		return nil, err
	}

	var releasedAt time.Time
	if p.ReleasedAt != nil {
		if t, ok := p.ReleasedAt.(time.Time); ok {
			releasedAt = t
		}
	}

	return &entity.Photograph{
		ID:           p.ID,
		Name:         p.Name,
		Abbreviation: p.Abbreviation,
		PhotoType:    p.PhotoType,
		GroupID:      p.GroupID,
		ReleasedAt:   releasedAt,
	}, nil
}

func (x *Photgraph) ListAllForAdmin(ctx context.Context) ([]entity.PhotographWithDetails, error) {
	ps, err := x.Querier.GetPhotographListForAdmin(ctx, x.DB)
	if err != nil {
		return nil, err
	}

	var photographs []entity.PhotographWithDetails
	for _, p := range ps {
		var releasedAt time.Time
		if p.ReleasedAt != nil {
			if t, ok := p.ReleasedAt.(time.Time); ok {
				releasedAt = t
			}
		}

		photograph := entity.PhotographWithDetails{
			ID:           p.ID,
			Name:         p.Name,
			Abbreviation: p.Abbreviation,
			PhotoType:    p.PhotoType,
			GroupID:      p.GroupID,
			ReleasedAt:   releasedAt,
			NameForOrder: p.NameForOrder,
			GroupName:    p.GroupName,
		}
		photographs = append(photographs, photograph)
	}
	return photographs, nil
}

func (x *Photgraph) Update(ctx context.Context, id int64, params UpdatePhotographParams) error {
	err := x.Querier.UpdatePhotograph(ctx, x.DB, repository.UpdatePhotographParams{
		ID:           id,
		Name:         params.Name,
		GroupID:      params.GroupID,
		PhotoType:    params.PhotoType,
		Abbreviation: params.Abbreviation,
		NameForOrder: params.NameForOrder,
	})
	if err != nil {
		return err
	}
	return nil
}

func (x *Photgraph) Delete(ctx context.Context, id int64) error {
	err := x.Querier.DeletePhotograph(ctx, x.DB, id)
	if err != nil {
		return err
	}
	return nil
}
