package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type ListScene struct {
	Service ListSceneService
}

func (ls *ListScene) ListScene(c *gin.Context) {
	ctx := context.Background()

	var req service.ListSceneRequest
	// bind request params to object
	c.ShouldBind(&req)

	if req.Color == "" {
		req.Color = "%"
	}
	if req.Member == "" {
		req.Member = "%"
	} else {
		req.Member = fmt.Sprintf("%%%s%%", req.Member)
	}

	if req.Photograph == "" {
		req.Photograph = "%"
	} else {
		req.Photograph = fmt.Sprintf("%%%s%%", req.Photograph)
	}
	req.FullName = true

	ss, err := ls.Service.ListScene(ctx, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "scenes/index.go.tmpl", gin.H{
		"title":  "Scenes Index",
		"scenes": ss,
	})
}
