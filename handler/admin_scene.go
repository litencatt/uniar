package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type AdminSceneHandler struct {
	SceneService *service.Scene
}

func (h *AdminSceneHandler) ListScene(c *gin.Context) {
	ctx := c.Request.Context()

	// ページネーション
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	perPage := int64(20)
	offset := (page - 1) * perPage

	scenes, err := h.SceneService.ListForAdmin(ctx, perPage, offset)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	total, err := h.SceneService.CountForAdmin(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	totalPages := (total + perPage - 1) / perPage
	hasNext := page < totalPages
	hasPrev := page > 1

	c.HTML(http.StatusOK, "admin/scene/list.go.tmpl", gin.H{
		"scenes":      scenes,
		"currentPage": page,
		"totalPages":  totalPages,
		"hasNext":     hasNext,
		"hasPrev":     hasPrev,
		"nextPage":    page + 1,
		"prevPage":    page - 1,
	})
}

func (h *AdminSceneHandler) ShowScene(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"error": "Invalid scene ID",
		})
		return
	}

	scene, err := h.SceneService.GetByID(ctx, id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404.go.tmpl", gin.H{
			"error": "Scene not found",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/scene/show.go.tmpl", gin.H{
		"scene": scene,
	})
}

func (h *AdminSceneHandler) EditScene(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"error": "Invalid scene ID",
		})
		return
	}

	scene, err := h.SceneService.GetByID(ctx, id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404.go.tmpl", gin.H{
			"error": "Scene not found",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/scene/edit.go.tmpl", gin.H{
		"scene": scene,
	})
}

func (h *AdminSceneHandler) UpdateScene(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
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
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
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
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"error": "Invalid scene ID",
		})
		return
	}

	err = h.SceneService.Delete(ctx, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/scene")
}