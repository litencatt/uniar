package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	goauth "google.golang.org/api/oauth2/v2"
)

var LoginInfo SessionInfo

type SessionInfo struct {
	UserId   interface{}
	UserMail interface{}
	LoggedIn bool
}

func TopHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "top/index.go.tmpl", gin.H{
		"title":    "Main Index",
		"LoggedIn": LoginInfo.LoggedIn,
		"EMail":    LoginInfo.UserMail,
	})
}

type LoginProducer struct {
	ProducerService ProducerService
	MemberService   MemberService
}

func (x *LoginProducer) AuthHandler(c *gin.Context) {
	ctx := context.Background()
	user, exits := c.Get("user")
	if exits {
		LoginInfo.UserId = user.(goauth.Userinfo).Id
		LoginInfo.UserMail = user.(goauth.Userinfo).Email
		LoginInfo.LoggedIn = true
		p, err := x.ProducerService.FindProducer(ctx, LoginInfo.UserId.(string))
		switch {
		case err == sql.ErrNoRows:
			fmt.Println("FindProducer record not found")
		case err != nil:
			fmt.Println("FindProducer error")
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		if p.IdentityId == "" {
			if err := x.ProducerService.RegistProducer(ctx, LoginInfo.UserId.(string), ""); err != nil {
				fmt.Println("UpdateProducer error")
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	ClearLoginInfo()
	fmt.Println("ログアウト完了")
	c.Redirect(http.StatusMovedPermanently, "/")
}

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exits := c.Get("user")
		if exits {
			LoginInfo.UserId = user.(goauth.Userinfo).Id
			LoginInfo.UserMail = user.(goauth.Userinfo).Email
			LoginInfo.LoggedIn = true
		}
	}
}

func ClearLoginInfo() {
	LoginInfo = SessionInfo{}
}
