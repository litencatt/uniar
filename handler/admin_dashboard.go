package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type AdminDashboardHandler struct {
	MusicService      *service.Music
	PhotographService *service.Photgraph
	SceneService      *service.Scene
}

func (h *AdminDashboardHandler) Dashboard(c *gin.Context) {
	ctx := c.Request.Context()

	// 各エンティティの統計情報を取得
	musics, err := h.MusicService.ListAll(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	photographs, err := h.PhotographService.ListAllForAdmin(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	sceneCount, err := h.SceneService.CountForAdmin(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	// 最新のシーンカードを数件取得
	recentScenes, err := h.SceneService.ListForAdmin(ctx, 5, 0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "admin/dashboard.go.tmpl", gin.H{
		"musicCount":      len(musics),
		"photographCount": len(photographs),
		"sceneCount":      sceneCount,
		"recentScenes":    recentScenes,
	})
}