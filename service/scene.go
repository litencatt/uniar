package service

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strconv"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

type Scene struct {
	DB      *sql.DB
	Querier repository.Querier
}

type ListSceneRequest struct {
	Color      string `form:"color"`
	Member     string `form:"member"`
	Photograph string `form:"photograph"`
	Sort       string `form:"sort"`
	Have       bool   `form:"have"`
	NotHave    bool   `form:"not_have"`
	Detail     bool   `form:"detail"`
	FullName   bool   `form:"full_name"`
}

type ListSceneAllRequest struct {
	Color      string `form:"color"`
	Member     string `form:"member"`
	Photograph string `form:"photograph"`
	Sort       string `form:"sort"`
	Have       bool   `form:"have"`
	NotHave    bool   `form:"not_have"`
	Detail     bool   `form:"detail"`
	FullName   bool   `form:"full_name"`
	GroupId    int64
	ProducerID int64
}

func (x *Scene) ListSceneAll(ctx context.Context, arg *ListSceneAllRequest) ([]entity.Scene, []entity.ProducerScene, error) {
	as, err := x.Querier.GetAllScenesWithGroupId(ctx, x.DB, arg.GroupId)
	if err != nil {
		return nil, nil, err
	}
	gs, err := x.Querier.GetScenesWithGroupId(ctx, x.DB, repository.GetScenesWithGroupIdParams{
		GroupID:    arg.GroupId,
		ProducerID: arg.ProducerID,
	})
	if err != nil {
		return nil, nil, err
	}

	var ss []entity.Scene
	for _, s := range as {
		scene := entity.Scene{
			PhotographID: s.PhotographID,
			MemberID:     s.MemberID,
			SsrPlus:      s.SsrPlus == 1,
		}
		ss = append(ss, scene)
	}

	var ps []entity.ProducerScene
	for _, s := range gs {
		scene := entity.ProducerScene{
			PhotographID: s.PhotographID,
			MemberID:     s.MemberID,
			SsrPlus:      s.SsrPlus == 1,
		}
		ps = append(ps, scene)
	}
	return ss, ps, nil
}

func (x *Scene) ListScene(ctx context.Context, arg *ListSceneRequest) ([]entity.Scene, error) {
	ss, err := x.Querier.GetScenesWithColor(ctx, x.DB, repository.GetScenesWithColorParams{
		Name:   arg.Color,
		Name_2: arg.Member,
		Name_3: arg.Photograph,
	})
	if err != nil {
		fmt.Println("GetScenesWithColor error.")
		return nil, err
	}

	var scenes []entity.Scene
	for _, s := range ss {
		// Show only scene you have
		if arg.Have && s.PsHave.(int64) == 0 {
			continue
		}

		// Show only scene you not have
		if arg.NotHave && s.PsHave.(int64) == 1 {
			continue
		}

		var e float64
		if s.ExpectedValue.Valid {
			e, _ = strconv.ParseFloat(s.ExpectedValue.String, 32)
		}
		p := s.Photograph
		if !arg.FullName && s.Abbreviation != "" {
			p = s.Abbreviation
		}
		scene := entity.Scene{
			Photograph: p,
			Member:     s.Member,
			Color:      s.Color,
			Total:      s.Total,
			Vo:         s.VocalMax,
			Da:         s.DanceMax,
			Pe:         s.PerformanceMax,
			Expect:     float32(e),
			SsrPlus:    s.SsrPlus == 1,
		}
		scene.CalcTotal(s.Bonds, s.Discography)
		scenes = append(scenes, scene)
	}

	// 各センタースキル毎の順位
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
	for i, _ := range scenes {
		scenes[i].All35 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoDa50Score > scenes[j].VoDa50Score })
	for i, _ := range scenes {
		scenes[i].VoDa50 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].DaPe50Score > scenes[j].DaPe50Score })
	for i, _ := range scenes {
		scenes[i].DaPe50 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoPe50Score > scenes[j].VoPe50Score })
	for i, _ := range scenes {
		scenes[i].VoPe50 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].Vo85Score > scenes[j].Vo85Score })
	for i, _ := range scenes {
		scenes[i].Vo85 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].Da85Score > scenes[j].Da85Score })
	for i, _ := range scenes {
		scenes[i].Da85 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].Pe85Score > scenes[j].Pe85Score })
	for i, _ := range scenes {
		scenes[i].Pe85 = int64(i + 1)
	}

	// 指定ソートで並び替え
	switch arg.Sort {
	case "All35":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
	case "VoDa50":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoDa50Score > scenes[j].VoDa50Score })
	case "DaPe50":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].DaPe50Score > scenes[j].DaPe50Score })
	case "VoPe50":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoPe50Score > scenes[j].VoPe50Score })
	case "Vo85":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Vo85Score > scenes[j].Vo85Score })
	case "Da85":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Da85Score > scenes[j].Da85Score })
	case "Pe85":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Pe85Score > scenes[j].Pe85Score })
	default:
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
	}

	return scenes, nil
}
