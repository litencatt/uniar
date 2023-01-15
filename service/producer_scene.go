package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

type ProducerScene struct {
	DB      *sql.DB
	Querier repository.Querier
}

type ListProducerSceneRequest struct {
	Color      string `form:"color"`
	Member     string `form:"member"`
	Photograph string `form:"photograph"`
	Sort       string `form:"sort"`
	Have       bool   `form:"have"`
	NotHave    bool   `form:"not_have"`
	Detail     bool   `form:"detail"`
	FullName   bool   `form:"full_name"`
}

func (x *ProducerScene) ListScene(ctx context.Context, arg *ListSceneRequest) ([]entity.Scene, error) {
	pss, err := x.Querier.GetProducerScenesByGroupId(ctx, x.DB, 1)
	if err != nil {
		return nil, err
	}

	var scenes []entity.Scene
	for _, s := range pss {
		scene := entity.Scene{
			Photograph: s.Photograph,
			Member:     s.Member,
			Color:      s.Color,
			Have:       s.Have == int64(1),
		}
		scenes = append(scenes, scene)
	}

	return scenes, nil
}

type RegistProducerSceneRequest struct {
	ProducerID   int64
	PhotographID int64
	MemberID     int64
}

func (x *ProducerScene) RegistScene(ctx context.Context, arg *RegistProducerSceneRequest) error {
	fmt.Printf("%+v\n", arg)
	if err := x.Querier.UpdateProducerScene(ctx, x.DB, repository.UpdateProducerSceneParams{
		ProducerID:   arg.ProducerID,
		PhotographID: arg.PhotographID,
		MemberID:     arg.MemberID,
		Have:         int64(1),
	}); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
