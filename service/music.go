package service

import (
	"context"
	"database/sql"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

type Music struct {
	DB      *sql.DB
	Querier repository.Querier
}

type UpdateMusicParams struct {
	Name        string
	Normal      int64
	Pro         int64
	Master      int64
	Length      int64
	ColorTypeID int64
	LiveID      int64
	MusicBonus  int64
}

func (s *Music) GetByID(ctx context.Context, id int64) (*entity.Music, error) {
	m, err := s.Querier.GetMusicById(ctx, s.DB, id)
	if err != nil {
		return nil, err
	}

	musicBonus := int64(0)
	if m.MusicBonus.Valid {
		musicBonus = m.MusicBonus.Int64
	}

	return &entity.Music{
		ID:          m.ID,
		Name:        m.Name,
		Normal:      m.Normal,
		Pro:         m.Pro,
		Master:      m.Master,
		Length:      m.Length,
		ColorTypeID: m.ColorTypeID,
		LiveID:      m.LiveID,
		MusicBonus:  musicBonus,
	}, nil
}

func (s *Music) ListAll(ctx context.Context) ([]entity.MusicWithDetails, error) {
	ms, err := s.Querier.GetMusicListAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var musics []entity.MusicWithDetails
	for _, m := range ms {
		musicBonus := int64(0)
		if m.MusicBonus.Valid {
			musicBonus = m.MusicBonus.Int64
		}

		music := entity.MusicWithDetails{
			ID:          m.ID,
			Name:        m.Name,
			Normal:      m.Normal,
			Pro:         m.Pro,
			Master:      m.Master,
			Length:      m.Length,
			ColorTypeID: m.ColorTypeID,
			LiveID:      m.LiveID,
			MusicBonus:  musicBonus,
			LiveName:    m.LiveName,
			ColorName:   m.ColorName,
		}
		musics = append(musics, music)
	}
	return musics, nil
}

func (s *Music) Update(ctx context.Context, id int64, params UpdateMusicParams) error {
	musicBonus := sql.NullInt64{
		Int64: params.MusicBonus,
		Valid: true,
	}

	err := s.Querier.UpdateMusic(ctx, s.DB, repository.UpdateMusicParams{
		ID:          id,
		Name:        params.Name,
		Normal:      params.Normal,
		Pro:         params.Pro,
		Master:      params.Master,
		Length:      params.Length,
		ColorTypeID: params.ColorTypeID,
		LiveID:      params.LiveID,
		MusicBonus:  musicBonus,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Music) Delete(ctx context.Context, id int64) error {
	err := s.Querier.DeleteMusic(ctx, s.DB, id)
	if err != nil {
		return err
	}
	return nil
}