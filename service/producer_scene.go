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
	ProducerId int64
}

func (x *ProducerScene) ListScene(ctx context.Context, arg *ListProducerSceneRequest) ([]entity.ProducerScene,  error) {
	ps, err := x.Querier.GetScenesWithGroupId(ctx, x.DB, repository.GetScenesWithGroupIdParams{
		GroupID: arg.GroupID,
		ProducerID: arg.ProducerId,
	})
	if err != nil {
		return nil, err
	}
	var producerScenes []entity.ProducerScene
	for _, s := range ps {
		scene := entity.ProducerScene{
			PhotographID: s.PhotographID,
			MemberID:     s.MemberID,
			SsrPlus:      s.SsrPlus == 1,
		}
		producerScenes = append(producerScenes, scene)
	}

	return producerScenes, nil
}

type RegistProducerSceneRequest struct {
	ProducerID   int64
	PhotographID int64
	MemberID     int64
	SsrPlus      int64
}

func (x *ProducerScene) RegistScene(ctx context.Context, arg *RegistProducerSceneRequest) error {
	// fmt.Printf("%+v\n", arg)
	if err := x.Querier.RegistProducerScene(ctx, x.DB, repository.RegistProducerSceneParams{
		ProducerID:   arg.ProducerID,
		PhotographID: arg.PhotographID,
		MemberID:     arg.MemberID,
		SsrPlus:      arg.SsrPlus,
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
