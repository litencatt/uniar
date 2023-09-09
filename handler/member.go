package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/entity"
)

type ListMember struct {
	SceneService      SceneService
	MemberService     MemberService
	PhotographService PhotographService
}

func (ls *ListMember) ListMember(c *gin.Context) {
	ctx := context.Background()
	fmt.Println("ListMember() start")
	fmt.Printf("User:%+v\n", User)

	ms, err := ls.MemberService.ListProducerMember(ctx, User.ProducerId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "members/index.go.tmpl", gin.H{
		"title":    "Members",
		"LoggedIn": User.LoggedIn,
		"EMail":    User.EMail,
		"members":  ms,
	})
}

func (ls *ListMember) UpdateMember(c *gin.Context) {
	ctx := context.Background()

	pms, err := ls.MemberService.ListProducerMember(ctx, User.ProducerId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Request.ParseForm()
	for _, pm := range pms {
		bond := c.Request.Form[fmt.Sprintf("bonds_%d", pm.MemberID)]
		bondInt, _ := strconv.ParseInt(bond[0], 10, 64)

		disco := c.Request.Form[fmt.Sprintf("disco_%d", pm.MemberID)]
		discoInt, _ := strconv.ParseInt(disco[0], 10, 64)

		ls.MemberService.UpdateProducerMember(ctx, entity.ProducerMember{
			MemberID:    pm.MemberID,
			BondLevel:   bondInt,
			Discography: discoInt,
		})
	}

	c.Redirect(http.StatusFound, "/auth/members")
}
