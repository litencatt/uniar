package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type RegistScene struct {
	SceneService      SceneService
	MemberService     MemberService
	PhotographService PhotographService
}

func (x *RegistScene) GetRegist(c *gin.Context) {
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

	sss, err := x.SceneService.ListScene(ctx, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	sms, err := x.MemberService.GetMemberByGroup(ctx, 1)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	sps, err := x.PhotographService.GetPhotographByGroup(ctx, 1)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "regist/index.go.tmpl", gin.H{
		"title":        "Regist Index",
		"s_photograph": sps,
		"s_member":     sms,
		"s_scenes":     sss,
	})
}

func (x *RegistScene) PostRegist(c *gin.Context) {
	ctx := context.Background()

	c.Request.ParseForm()
	sakuraMembers, err := x.MemberService.GetMemberByGroup(ctx, 1)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	for _, m := range sakuraMembers {
		pids := c.Request.Form[fmt.Sprintf("member_%d[]", m.ID)]
		for _, pid := range pids {
			id, _ := strconv.ParseInt(pid, 10, 64)
			if err := x.SceneService.RegistScene(ctx, &service.RegistSceneRequest{
				ProducerID:   1,
				PhotographID: id,
				MemberID:     m.ID,
			}); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	c.HTML(http.StatusCreated, "regist/index.go.tmpl", gin.H{
		"title": "Regist Index",
	})
}
