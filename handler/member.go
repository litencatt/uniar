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
	us, _ := getUserSession(c)

	ms, err := ls.MemberService.ListProducerMember(ctx, us.ProducerId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "members/index.go.tmpl", gin.H{
		"title":    "Members",
		"LoggedIn": us.LoggedIn,
		"EMail":    us.EMail,
		"IsAdmin":  us.IsAdmin,
		"members":  ms,
	})
}

func (ls *ListMember) UpdateMember(c *gin.Context) {
	ctx := context.Background()
	fmt.Println("UpdateMember() start")
	us, _ := getUserSession(c)

	pms, err := ls.MemberService.ListProducerMember(ctx, us.ProducerId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.Request.ParseForm(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	for _, pm := range pms {
		bond := c.Request.Form[fmt.Sprintf("bonds_%d", pm.MemberID)]
		bondInt, _ := strconv.ParseInt(bond[0], 10, 64)

		disco := c.Request.Form[fmt.Sprintf("disco_%d", pm.MemberID)]
		discoInt, _ := strconv.ParseInt(disco[0], 10, 64)

		if err := ls.MemberService.UpdateProducerMember(ctx, entity.ProducerMember{
			ProducerID:  us.ProducerId,
			MemberID:    pm.MemberID,
			BondLevel:   bondInt,
			Discography: discoInt,
		}); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Redirect(http.StatusFound, "/auth/members")
}
