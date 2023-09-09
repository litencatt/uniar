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

	var group Group
	c.ShouldBindUri(&group)
	groupId, err := strconv.ParseInt(group.GroupID, 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	sceneList, err := x.SceneService.ListSceneAll(ctx, &service.ListSceneAllRequest{
		Photograph: "%",
		Color:      "%",
		Member:     "%",
		FullName:   true,
		GroupId:    groupId,
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
	for _, s := range sceneList {
		if s.SsrPlus {
			continue
		}
		producerScenes[s.PhotographID][s.MemberID] = s.Have
	}

	c.HTML(http.StatusOK, "regist/index.go.tmpl", gin.H{
		"title":          "Regist Index",
		"LoggedIn":       User.LoggedIn,
		"EMail":          User.EMail,
		"groupId":        groupId,
		"photos":         photos,
		"members":        members,
		"producerScenes": producerScenes,
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

	members, err := x.MemberService.GetMemberByGroup(ctx, groupId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	ssrPlusPhotographList, err := x.PhotographService.GetSsrPlusReleasedPhotographList(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var ssrPlusPhotographIDs []int
	for _, sp := range ssrPlusPhotographList {
		ssrPlusPhotographIDs = append(ssrPlusPhotographIDs, int(sp.ID))
	}

	c.Request.ParseForm()
	for _, m := range members {
		// Update ps.Have = 0
		if err := x.ProducerSceneService.InitAllScene(ctx, &service.InitProducerSceneRequest{
			ProducerID: 1,
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
				ProducerID:   1,
				PhotographID: photoId,
				MemberID:     m.ID,
				SsrPlus:      int64(0),
				Have:         int64(1),
			}); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			if include(ssrPlusPhotographIDs, int(photoId)) {
				if err := x.ProducerSceneService.RegistScene(ctx, &service.RegistProducerSceneRequest{
					ProducerID:   1,
					PhotographID: photoId,
					MemberID:     m.ID,
					SsrPlus:      int64(1),
					Have:         int64(1),
				}); err != nil {
					c.String(http.StatusInternalServerError, err.Error())
					return
				}
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
