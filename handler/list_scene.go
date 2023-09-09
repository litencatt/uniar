package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type ListScene struct {
	SceneService      SceneService
	MemberService     MemberService
	PhotographService PhotographService
}

func (ls *ListScene) ListScene(c *gin.Context) {
	ctx := context.Background()
	fmt.Println("ListScene() start")
	fmt.Printf("User:%+v\n", User)

	var req service.ListSceneRequest
	// bind request params to object
	c.ShouldBind(&req)
	// fmt.Printf("%+v\n", req)
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
	req.ProducerID = User.ProducerId

	ss, err := ls.SceneService.ListScene(ctx, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	ms, err := ls.MemberService.ListMember(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	ps, err := ls.PhotographService.ListPhotograph(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "scenes/index.go.tmpl", gin.H{
		"title":              "Scenes Index",
		"LoggedIn":           User.LoggedIn,
		"EMail":              User.EMail,
		"photograph":         ps,
		"selectedPhotograph": req.Photograph,
		"color":              []string{"Red", "Blue", "Green", "Yellow", "Purple"},
		"selectedColor":      req.Color,
		"member":             ms,
		"selectedMember":     req.Member,
		"have":               req.Have,
		"notHave":            req.NotHave,
		"detail":             req.Detail,
		"sort":               []string{"All35", "VoDa50", "DaPe50", "VoPe50", "Vo85", "Da85", "Pe85"},
		"selectedSort":       req.Sort,
		"scenes":             ss,
	})
}
