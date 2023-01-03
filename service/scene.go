package service

import (
	"context"
	"database/sql"
	"sort"
	"strconv"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

type ListScene struct {
	DB      *sql.DB
	Querier repository.Querier
}

type ListSceneRequest struct {
	Color      string
	Member     string
	Photograph string
	Sort       string
	Have       bool
	NotHave    bool
	Detail     bool
	FullName   bool
}

func (x *ListScene) ListScene(ctx context.Context, arg *ListSceneRequest) ([]entity.Scene, error) {

	ss, err := x.Querier.GetScenesWithColor(ctx, x.DB, repository.GetScenesWithColorParams{
		Name:   arg.Color,
		Name_2: arg.Member,
		Name_3: arg.Photograph,
	})
	if err != nil {
		return nil, err
	}

	var scenes []entity.Scene
	for _, s := range ss {
		// Show only scene you have
		if arg.Have && s.Have == 0 {
			continue
		}
		// Show only scene you not have
		if arg.NotHave && s.Have != 0 {
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
	case "all35":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
	case "voda50":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoDa50Score > scenes[j].VoDa50Score })
	case "dape50":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].DaPe50Score > scenes[j].DaPe50Score })
	case "vope50":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoPe50Score > scenes[j].VoPe50Score })
	case "vo85":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Vo85Score > scenes[j].Vo85Score })
	case "da85":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Da85Score > scenes[j].Da85Score })
	case "pe85":
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Pe85Score > scenes[j].Pe85Score })
	default:
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
	}

	return scenes, nil
}
