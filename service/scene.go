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
	ProducerID int64
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

	ps, err := x.Querier.GetProducerScenesWithProducerId(ctx, x.DB, arg.ProducerID)
	if err != nil {
		fmt.Println("GetProducerScenesWithProducerId error.")
		return nil, err
	}
	//fmt.Printf("%+v\n", ps)

	var scenes []entity.Scene
	for _, s := range ss {
		have := false
		for _, p := range ps {
			if s.PhotographID == p.PhotographID && s.MemberID == p.MemberID {
				have = true
				break
			}
		}
		// Show only scene you have
		if arg.Have && !have {
			continue
		}

		// Show only scene you not have
		if arg.NotHave && have {
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
		bonds := int64(0)
		if s.Bonds.Valid {
			bonds = s.Bonds.Int64
		}
		discography := int64(0)
		if s.Discography.Valid {
			discography = s.Discography.Int64
		}
		scene.CalcTotal(bonds, discography)
		scenes = append(scenes, scene)
	}

	// 各センタースキル毎の順位
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
	for i := range scenes {
		scenes[i].All35 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoDa50Score > scenes[j].VoDa50Score })
	for i := range scenes {
		scenes[i].VoDa50 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].DaPe50Score > scenes[j].DaPe50Score })
	for i := range scenes {
		scenes[i].DaPe50 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoPe50Score > scenes[j].VoPe50Score })
	for i := range scenes {
		scenes[i].VoPe50 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].Vo85Score > scenes[j].Vo85Score })
	for i := range scenes {
		scenes[i].Vo85 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].Da85Score > scenes[j].Da85Score })
	for i := range scenes {
		scenes[i].Da85 = int64(i + 1)
	}
	sort.Slice(scenes, func(i, j int) bool { return scenes[i].Pe85Score > scenes[j].Pe85Score })
	for i := range scenes {
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

type UpdateSceneParams struct {
	PhotographID   int64
	MemberID       int64
	ColorTypeID    int64
	VocalMax       int64
	DanceMax       int64
	PerformanceMax int64
	CenterSkill    string
	ExpectedValue  string
	SsrPlus        int64
}

type AddSceneParams struct {
	PhotographID   int64
	MemberID       int64
	ColorTypeID    int64
	VocalMax       int64
	DanceMax       int64
	PerformanceMax int64
	CenterSkill    string
	ExpectedValue  string
	SsrPlus        int64
}

func (x *Scene) GetByID(ctx context.Context, id int64) (*entity.SceneWithDetails, error) {
	s, err := x.Querier.GetSceneById(ctx, x.DB, id)
	if err != nil {
		return nil, err
	}

	centerSkill := ""
	if s.CenterSkill.Valid {
		centerSkill = s.CenterSkill.String
	}

	expectedValue := ""
	if s.ExpectedValue.Valid {
		expectedValue = s.ExpectedValue.String
	}

	return &entity.SceneWithDetails{
		ID:             s.ID,
		PhotographID:   s.PhotographID,
		MemberID:       s.MemberID,
		ColorTypeID:    s.ColorTypeID,
		VocalMax:       s.VocalMax,
		DanceMax:       s.DanceMax,
		PerformanceMax: s.PerformanceMax,
		CenterSkill:    centerSkill,
		ExpectedValue:  expectedValue,
		SsrPlus:        s.SsrPlus,
		PhotographName: s.PhotographName,
		MemberName:     s.MemberName,
		ColorName:      s.ColorName,
	}, nil
}

func (x *Scene) ListForAdmin(ctx context.Context, limit, offset int64) ([]entity.SceneWithDetails, error) {
	ss, err := x.Querier.GetSceneListForAdmin(ctx, x.DB, repository.GetSceneListForAdminParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var scenes []entity.SceneWithDetails
	for _, s := range ss {
		centerSkill := ""
		if s.CenterSkill.Valid {
			centerSkill = s.CenterSkill.String
		}

		expectedValue := ""
		if s.ExpectedValue.Valid {
			expectedValue = s.ExpectedValue.String
		}

		scene := entity.SceneWithDetails{
			ID:             s.ID,
			PhotographID:   s.PhotographID,
			MemberID:       s.MemberID,
			ColorTypeID:    s.ColorTypeID,
			VocalMax:       s.VocalMax,
			DanceMax:       s.DanceMax,
			PerformanceMax: s.PerformanceMax,
			CenterSkill:    centerSkill,
			ExpectedValue:  expectedValue,
			SsrPlus:        s.SsrPlus,
			PhotographName: s.PhotographName,
			MemberName:     s.MemberName,
			ColorName:      s.ColorName,
		}
		scenes = append(scenes, scene)
	}
	return scenes, nil
}

func (x *Scene) CountForAdmin(ctx context.Context) (int64, error) {
	count, err := x.Querier.CountScenesForAdmin(ctx, x.DB)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (x *Scene) Update(ctx context.Context, id int64, params UpdateSceneParams) error {
	centerSkill := sql.NullString{
		String: params.CenterSkill,
		Valid:  params.CenterSkill != "",
	}

	expectedValue := sql.NullString{
		String: params.ExpectedValue,
		Valid:  params.ExpectedValue != "",
	}

	err := x.Querier.UpdateScene(ctx, x.DB, repository.UpdateSceneParams{
		ID:             id,
		PhotographID:   params.PhotographID,
		MemberID:       params.MemberID,
		ColorTypeID:    params.ColorTypeID,
		VocalMax:       params.VocalMax,
		DanceMax:       params.DanceMax,
		PerformanceMax: params.PerformanceMax,
		CenterSkill:    centerSkill,
		ExpectedValue:  expectedValue,
		SsrPlus:        params.SsrPlus,
	})
	if err != nil {
		return err
	}
	return nil
}

func (x *Scene) Delete(ctx context.Context, id int64) error {
	err := x.Querier.DeleteScene(ctx, x.DB, id)
	if err != nil {
		return err
	}
	return nil
}

func (x *Scene) Add(ctx context.Context, params AddSceneParams) error {
	centerSkill := sql.NullString{
		String: params.CenterSkill,
		Valid:  params.CenterSkill != "",
	}

	expectedValue := sql.NullString{
		String: params.ExpectedValue,
		Valid:  params.ExpectedValue != "",
	}

	err := x.Querier.RegistScene(ctx, x.DB, repository.RegistSceneParams{
		PhotographID:   params.PhotographID,
		MemberID:       params.MemberID,
		ColorTypeID:    params.ColorTypeID,
		VocalMax:       params.VocalMax,
		DanceMax:       params.DanceMax,
		PerformanceMax: params.PerformanceMax,
		CenterSkill:    centerSkill,
		ExpectedValue:  expectedValue,
		SsrPlus:        params.SsrPlus,
	})
	if err != nil {
		return err
	}
	return nil
}
