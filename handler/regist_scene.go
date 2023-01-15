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

func (x *RegistScene) GetRegist(c *gin.Context) {
	ctx := context.Background()
	sps, err := x.ProducerSceneService.ListScene(ctx, &service.ListSceneRequest{
		Photograph: "%",
		Color:      "%",
		Member:     "%",
		FullName:   true,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	sakuraMembers, err := x.MemberService.GetMemberByGroup(ctx, 1)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	sakuraPhotos, err := x.PhotographService.GetPhotographByGroup(ctx, 1)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	sakuraProducerScenes := make([][]int64, 120)
	fmt.Println(len(sakuraProducerScenes))
	for i := 0; i < 120; i++ {
		sakuraProducerScenes[i] = make([]int64, 100)
	}
	for _, ps := range sps {
		//fmt.Printf("%+v\n", ps)
		sakuraProducerScenes[ps.PhotographID-1][ps.MemberID-1] = ps.Have
	}
	fmt.Println(sakuraProducerScenes)

	c.HTML(http.StatusOK, "regist/index.go.tmpl", gin.H{
		"title":                "Regist Index",
		"sakuraPhotos":         sakuraPhotos,
		"sakuraMembers":        sakuraMembers,
		"sakuraProducerScenes": sakuraProducerScenes,
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
		fmt.Println(pids)
		for _, pid := range pids {
			id, _ := strconv.ParseInt(pid, 10, 64)
			if err := x.ProducerSceneService.RegistScene(ctx, &service.RegistProducerSceneRequest{
				ProducerID:   1,
				PhotographID: id,
				MemberID:     m.ID,
			}); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
	c.Redirect(http.StatusFound, "/regist")
}
