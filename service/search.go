package service

import (
	"context"
	"database/sql"

	"github.com/litencatt/uniar/repository"
)

type SearchService struct {
	DB      *sql.DB
	Querier repository.Querier
}

type MusicSearchParams struct {
	Name        string `form:"name"`
	LiveID      int64  `form:"live_id"`
	ColorTypeID int64  `form:"color_type_id"`
}

type PhotographSearchParams struct {
	Name      string `form:"name"`
	GroupID   int64  `form:"group_id"`
	PhotoType string `form:"photo_type"`
}

type SceneSearchParams struct {
	MemberID      int64 `form:"member_id"`
	PhotographID  int64 `form:"photograph_id"`
	ColorTypeID   int64 `form:"color_type_id"`
	SSRPlus       int   `form:"ssr_plus"`
}

func (s *SearchService) SearchMusic(ctx context.Context, params MusicSearchParams) ([]repository.SearchMusicListRow, error) {
	return s.Querier.SearchMusicList(ctx, s.DB, repository.SearchMusicListParams{
		Column1: params.Name,
		Column2: params.LiveID,
		Column3: params.ColorTypeID,
	})
}

func (s *SearchService) SearchPhotograph(ctx context.Context, params PhotographSearchParams) ([]repository.SearchPhotographListRow, error) {
	return s.Querier.SearchPhotographList(ctx, s.DB, repository.SearchPhotographListParams{
		Column1: params.Name,
		Column2: params.GroupID,
		Column3: params.PhotoType,
	})
}

func (s *SearchService) SearchScene(ctx context.Context, params SceneSearchParams) ([]repository.SearchSceneListRow, error) {
	return s.Querier.SearchSceneList(ctx, s.DB, repository.SearchSceneListParams{
		Column1: params.MemberID,
		Column2: params.PhotographID,
		Column3: params.ColorTypeID,
		Column4: int64(params.SSRPlus),
	})
}