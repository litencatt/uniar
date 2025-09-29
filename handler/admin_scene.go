package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/repository"
	"github.com/litencatt/uniar/service"
)

type AdminSceneHandler struct {
	SceneService  *service.Scene
	SearchService *service.SearchService
}

func (h *AdminSceneHandler) ListScene(c *gin.Context) {
	ctx := c.Request.Context()

	// ドロップダウン選択肢用のデータを取得
	members, err := h.SearchService.Querier.GetAllMembers(ctx, h.SearchService.DB)
	if err != nil {
		members = []repository.Member{} // エラー時は空配列
	}

	photographs, err := h.SearchService.Querier.GetPhotographListForAdmin(ctx, h.SearchService.DB)
	if err != nil {
		photographs = []repository.GetPhotographListForAdminRow{} // エラー時は空配列
	}

	colors, err := h.SearchService.Querier.GetColorTypeList(ctx, h.SearchService.DB)
	if err != nil {
		colors = []repository.GetColorTypeListRow{} // エラー時は空配列
	}

	// 検索パラメータを取得
	var searchParams service.SceneSearchParams
	if err := c.ShouldBindQuery(&searchParams); err == nil {
		// 検索パラメータがある場合は検索実行
		if searchParams.MemberID != 0 || searchParams.PhotographID != 0 || searchParams.ColorTypeID != 0 || searchParams.SSRPlus != -1 {
			scenes, err := h.SearchService.SearchScene(ctx, searchParams)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
					"error": err.Error(),
				})
				return
			}

			c.HTML(http.StatusOK, "admin_scene_list.go.tmpl", gin.H{
				"title":        "シーンカード管理",
				"scenes":       scenes,
				"searchParams": searchParams,
				"members":      members,
				"photographs":  photographs,
				"colors":       colors,
				// ページネーション情報は検索時は無効化
				"currentPage": 1,
				"totalPages":  1,
				"hasNext":     false,
				"hasPrev":     false,
			})
			return
		}
	}

	// 検索パラメータがない場合は従来のページネーション付きリスト
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	perPage := int64(20)
	offset := (page - 1) * perPage

	scenes, err := h.SceneService.ListForAdmin(ctx, perPage, offset)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	total, err := h.SceneService.CountForAdmin(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	totalPages := (total + perPage - 1) / perPage
	hasNext := page < totalPages
	hasPrev := page > 1

	c.HTML(http.StatusOK, "admin_scene_list.go.tmpl", gin.H{
		"title":        "シーンカード管理",
		"scenes":       scenes,
		"searchParams": service.SceneSearchParams{SSRPlus: -1}, // デフォルト値
		"members":      members,
		"photographs":  photographs,
		"colors":       colors,
		"currentPage":  page,
		"totalPages":   totalPages,
		"hasNext":      hasNext,
		"hasPrev":      hasPrev,
		"nextPage":     page + 1,
		"prevPage":     page - 1,
	})
}


func (h *AdminSceneHandler) EditScene(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid scene ID",
		})
		return
	}

	scene, err := h.SceneService.GetByID(ctx, id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.go.tmpl", gin.H{
			"error": "Scene not found",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/scene_edit.go.tmpl", gin.H{
		"scene": scene,
	})
}

func (h *AdminSceneHandler) UpdateScene(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid scene ID",
		})
		return
	}

	photographID, _ := strconv.ParseInt(c.PostForm("photograph_id"), 10, 64)
	memberID, _ := strconv.ParseInt(c.PostForm("member_id"), 10, 64)
	colorTypeID, _ := strconv.ParseInt(c.PostForm("color_type_id"), 10, 64)
	vocalMax, _ := strconv.ParseInt(c.PostForm("vocal_max"), 10, 64)
	danceMax, _ := strconv.ParseInt(c.PostForm("dance_max"), 10, 64)
	performanceMax, _ := strconv.ParseInt(c.PostForm("performance_max"), 10, 64)
	centerSkill := c.PostForm("center_skill")
	expectedValue := c.PostForm("expected_value")
	ssrPlus, _ := strconv.ParseInt(c.PostForm("ssr_plus"), 10, 64)

	params := service.UpdateSceneParams{
		PhotographID:   photographID,
		MemberID:       memberID,
		ColorTypeID:    colorTypeID,
		VocalMax:       vocalMax,
		DanceMax:       danceMax,
		PerformanceMax: performanceMax,
		CenterSkill:    centerSkill,
		ExpectedValue:  expectedValue,
		SsrPlus:        ssrPlus,
	}

	err = h.SceneService.Update(ctx, id, params)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/scene")
}

func (h *AdminSceneHandler) DeleteScene(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid scene ID",
		})
		return
	}

	err = h.SceneService.Delete(ctx, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/scene")
}

func (h *AdminSceneHandler) NewScene(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/scene_new.go.tmpl", gin.H{})
}

func (h *AdminSceneHandler) AddScene(c *gin.Context) {
	ctx := c.Request.Context()

	photographID, _ := strconv.ParseInt(c.PostForm("photograph_id"), 10, 64)
	memberID, _ := strconv.ParseInt(c.PostForm("member_id"), 10, 64)
	colorTypeID, _ := strconv.ParseInt(c.PostForm("color_type_id"), 10, 64)
	vocalMax, _ := strconv.ParseInt(c.PostForm("vocal_max"), 10, 64)
	danceMax, _ := strconv.ParseInt(c.PostForm("dance_max"), 10, 64)
	performanceMax, _ := strconv.ParseInt(c.PostForm("performance_max"), 10, 64)
	centerSkill := c.PostForm("center_skill")
	expectedValue := c.PostForm("expected_value")
	ssrPlus, _ := strconv.ParseInt(c.PostForm("ssr_plus"), 10, 64)

	params := service.AddSceneParams{
		PhotographID:   photographID,
		MemberID:       memberID,
		ColorTypeID:    colorTypeID,
		VocalMax:       vocalMax,
		DanceMax:       danceMax,
		PerformanceMax: performanceMax,
		CenterSkill:    centerSkill,
		ExpectedValue:  expectedValue,
		SsrPlus:        ssrPlus,
	}

	err := h.SceneService.Add(ctx, params)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/scene")
}