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
	SceneService         SceneService
	ProducerSceneService ProducerSceneService
	MemberService        MemberService
	PhotographService    PhotographService
}

type Group struct {
	GroupID string `uri:"group_id"`
}

func (x *RegistScene) GetRegist(c *gin.Context) {
	ctx := context.Background()
	fmt.Println("GetRegist() start")
	us, err := getUserSession(c)
	if err != nil {
		fmt.Printf("GetRegist() user session not found: %+v\n", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var group Group
	c.ShouldBindUri(&group)
	groupId, err := strconv.ParseInt(group.GroupID, 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	ss, ps, err := x.SceneService.ListSceneAll(ctx, &service.ListSceneAllRequest{
		Photograph: "%",
		Color:      "%",
		Member:     "%",
		FullName:   true,
		GroupId:    groupId,
		ProducerID: us.ProducerId,
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
		for j := 0; j < 100; j++ {
			producerScenes[i][j] = -1
		}
	}
	for _, s := range ss {
		if s.SsrPlus {
			continue
		}
		producerScenes[s.PhotographID][s.MemberID] = 0
	}
	// CheckBox ON if ProducerScene record exists
	for _, s := range ps {
		if s.SsrPlus {
			continue
		}
		producerScenes[s.PhotographID][s.MemberID] = 1
	}
	c.HTML(http.StatusOK, "regist/index.go.tmpl", gin.H{
		"title":          "Regist Index",
		"LoggedIn":       us.LoggedIn,
		"EMail":          us.EMail,
		"groupId":        groupId,
		"photos":         photos,
		"members":        members,
		"producerScenes": producerScenes,
	})
}

func (x *RegistScene) PostRegist(c *gin.Context) {
	ctx := context.Background()
	fmt.Println("PostRegist() start")
	us, err := getUserSession(c)
	if err != nil {
		fmt.Printf("PostRegist() user session not found: %+v\n", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var group Group
	c.ShouldBindUri(&group)
	groupId, err := strconv.ParseInt(group.GroupID, 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	members, err := x.MemberService.GetMemberByGroup(ctx, groupId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Request.ParseForm()
	// Delete a member's all producer_scenes once.
	// And then, insert only checkbox ON producer_scenes.
	for _, m := range members {
		if err := x.ProducerSceneService.InitAllScene(ctx, &service.InitProducerSceneRequest{
			ProducerID: us.ProducerId,
			MemberID:   m.ID,
		}); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// POST from form
		//   member_$member_id[]: $photograph_id
		// e.g. m.ID = 1
		//   member_1[]: 1
		//   member_1[]: 2
		//   ..
		// pids = ["1", "2"]
		photographIDs := c.Request.Form[fmt.Sprintf("member_%d[]", m.ID)]
		for _, pid := range photographIDs {
			photoId, _ := strconv.ParseInt(pid, 10, 64)
			if err := x.ProducerSceneService.RegistScene(ctx, &service.RegistProducerSceneRequest{
				ProducerID:   us.ProducerId,
				PhotographID: photoId,
				MemberID:     m.ID,
				SsrPlus:      int64(0),
			}); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/auth/regist/%d", groupId))
}

func include(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}
