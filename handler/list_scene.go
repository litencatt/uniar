package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type ListScene struct {
	Service ListSceneService
}

func (ls *ListScene) ListScene(c *gin.Context) {
	ctx := context.Background()
	ss, err := ls.Service.ListScene(ctx, &service.ListSceneRequest{
		Color:      "%",
		Member:     "%",
		Photograph: "キュン",
	})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "scenes/index.go.tmpl", gin.H{
		"title":  "Scenes Index",
		"scenes": ss,
	})
}
