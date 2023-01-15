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
	GroupID    int64
}

func (x *ProducerScene) ListScene(ctx context.Context, arg *ListProducerSceneRequest) ([]entity.ProducerScene, error) {
	pss, err := x.Querier.GetProducerScenesByGroupId(ctx, x.DB, arg.GroupID)
	if err != nil {
		return nil, err
	}

	var scenes []entity.ProducerScene
	for _, s := range pss {
		scene := entity.ProducerScene{
			PhotographID: s.PhotographID,
			MemberID:     s.MemberID,
			Have:         s.Have,
		}
		scenes = append(scenes, scene)
	}

	return scenes, nil
}

type RegistProducerSceneRequest struct {
	ProducerID   int64
	PhotographID int64
	MemberID     int64
	Have         int64
}

func (x *ProducerScene) RegistScene(ctx context.Context, arg *RegistProducerSceneRequest) error {
	fmt.Printf("%+v\n", arg)
	if err := x.Querier.UpdateProducerScene(ctx, x.DB, repository.UpdateProducerSceneParams{
		ProducerID:   arg.ProducerID,
		PhotographID: arg.PhotographID,
		MemberID:     arg.MemberID,
		Have:         arg.Have,
	}); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

type InitProducerSceneRequest struct {
	ProducerID int64
	MemberID   int64
}

func (x *ProducerScene) InitAllScene(ctx context.Context, arg *InitProducerSceneRequest) error {
	if err := x.Querier.InitProducerSceneAll(ctx, x.DB, repository.InitProducerSceneAllParams{
		ProducerID: arg.ProducerID,
		MemberID:   arg.MemberID,
	}); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
