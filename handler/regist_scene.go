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
	ProducerSceneService ProducerSceneService
	MemberService        MemberService
	PhotographService    PhotographService
}

type Group struct {
	GroupID string `uri:"group_id"`
}

func (x *RegistScene) GetRegist(c *gin.Context) {
	ctx := context.Background()

	var group Group
	c.ShouldBindUri(&group)
	groupId, err := strconv.ParseInt(group.GroupID, 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	sps, err := x.ProducerSceneService.ListScene(ctx, &service.ListProducerSceneRequest{
		Photograph: "%",
		Color:      "%",
		Member:     "%",
		FullName:   true,
		GroupID:    groupId,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	members, err := x.MemberService.GetMemberByGroup(ctx, groupId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	photos, err := x.PhotographService.GetPhotographByGroup(ctx, groupId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	producerScenes := make([][]int64, 120)
	for i := 0; i < 120; i++ {
		producerScenes[i] = make([]int64, 100)
	}
	for _, ps := range sps {
		producerScenes[ps.PhotographID][ps.MemberID] = ps.Have
	}

	c.HTML(http.StatusOK, "regist/index.go.tmpl", gin.H{
		"title":          "Regist Index",
		"photos":         photos,
		"members":        members,
		"producerScenes": producerScenes,
		"groupId":        groupId,
	})
}

func (x *RegistScene) PostRegist(c *gin.Context) {
	ctx := context.Background()

	var group Group
	c.ShouldBindUri(&group)
	groupId, err := strconv.ParseInt(group.GroupID, 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Request.ParseForm()
	members, err := x.MemberService.GetMemberByGroup(ctx, groupId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	for _, m := range members {
		if err := x.ProducerSceneService.InitAllScene(ctx, &service.InitProducerSceneRequest{
			ProducerID: 1,
			MemberID:   m.ID,
		}); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		pids := c.Request.Form[fmt.Sprintf("member_%d[]", m.ID)]
		// fmt.Println(pids)
		for _, pid := range pids {
			id, _ := strconv.ParseInt(pid, 10, 64)
			if err := x.ProducerSceneService.RegistScene(ctx, &service.RegistProducerSceneRequest{
				ProducerID:   1,
				PhotographID: id,
				MemberID:     m.ID,
				Have:         1,
			}); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/regist/%d", groupId))
}
