package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type RegistScene struct {
	SceneService ListSceneService
}

func (x *RegistScene) RegistScene(c *gin.Context) {
	ctx := context.Background()

	var req service.ListSceneRequest
	// bind request params to object
	c.ShouldBind(&req)
	fmt.Printf("%+v\n", req)
	if req.Photograph == "" {
		req.Photograph = "%"
	}
	if req.Color == "" {
		req.Color = "%"
	}
	if req.Member == "" {
		req.Member = "%"
	}
	req.FullName = true

	ss, err := x.SceneService.ListScene(ctx, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "regist/index.go.tmpl", gin.H{
		"title":  "Regist Index",
		"scenes": ss,
	})
}
