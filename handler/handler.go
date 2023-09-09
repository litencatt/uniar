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

var User UserSession

type UserSession struct {
	ProducerId       int64
	IdentityId	   string
	EMail    string
	LoggedIn bool
}

type LoginProducer struct {
	ProducerService ProducerService
	MemberService   MemberService
}

func RootHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
}

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("LoginCheck() start User:%+v\n", User)
		if isLoggedIn() {
			fmt.Println("LoginCheck() is logged in")
			//c.Redirect(http.StatusFound, "/auth")
		} else {
			fmt.Println("LoginCheck() is NOT logged in")
		}
	}
}

// AuthCheck is middleware for checking login status.
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("AuthCheck() start User:%+v\n", User)
		if isLoggedIn() {
			fmt.Println("AuthCheck() is logged in")
		} else {
			fmt.Println("AuthCheck() is NOT logged in")
			c.Redirect(http.StatusFound, "/auth")
		}
	}
}

func (x *LoginProducer) AuthHandler(c *gin.Context) {
	ctx := context.Background()
	session := sessions.Default(c)
	fmt.Println("AuthHandler() start")

	// ログイン済みならトップにリダイレクト
	sessionUser := session.Get("uniar_oauth_user")
	if sessionUser != nil {
		if sessionUser.(UserSession).LoggedIn {
			fmt.Println("already Loggedin. redirect to /auth/members")
			c.Redirect(http.StatusMovedPermanently, "/auth/members")
			c.Abort()
		}
	} 

	// contextに保存されたGoogle認証情報を取得
	ctxUser := c.MustGet("user")
	if user, ok := ctxUser.(goauth.Userinfo); ok {
		p, err := x.ProducerService.FindProducer(ctx, User.IdentityId)
		switch {
		case err == sql.ErrNoRows:
			fmt.Printf("producer not found. User.Id = %v\n", User.IdentityId)
		case err != nil:
			fmt.Printf("find producer error: User.Id = %v err = %v\n", User.IdentityId, err)
			c.AbortWithError(http.StatusInternalServerError, err)
		default:
			fmt.Printf("producer record found. User.Id = %v\n", User.IdentityId)
		}
		User.IdentityId = user.Id
		User.EMail = user.Email
		User.LoggedIn = true

		if p.IdentityId == "" {
			if err := x.ProducerService.RegistProducer(ctx, User.IdentityId); err != nil {
				fmt.Printf("update producer error. User.Id = %v\n", User.IdentityId)
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			p, err := x.ProducerService.FindProducer(ctx, User.IdentityId)
			switch {
			case err == sql.ErrNoRows:
				fmt.Printf("producer not found. User.Id = %v\n", User.IdentityId)
			case err != nil:
				fmt.Printf("find producer error: User.Id = %v err = %v\n", User.IdentityId, err)
				c.AbortWithError(http.StatusInternalServerError, err)
			default:
				fmt.Printf("producer record found. User.Id = %v\n", User.IdentityId)
			}
			User.ProducerId = p.ID
		} else {
			User.ProducerId = p.ID
		}
		fmt.Printf("User:%+v\n", User)

		session.Set("uniar_oauth_user", User)
		session.Save()
		ts := session.Get("uniar_oauth_user")
		fmt.Printf("key:uniar_oauth_user value: %+v\n", ts)
	} else {
		fmt.Println("Not Authorized")
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	c.Redirect(http.StatusMovedPermanently, "/auth/members")
}

func LogoutHandler(c *gin.Context) {
	c.Set("user", "")
	session := sessions.Default(c)
	// session key name of zalando/gin-oauth2
	// https://github.com/zalando/gin-oauth2/blob/master/google/google.go
	session.Set("ginoauth_google_session", "")
	session.Set("uniar_oauth_user", "")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Save()
	ClearUser()

	c.Redirect(http.StatusMovedPermanently, "/")
}

func isLoggedIn() bool {
	return User.LoggedIn
}

func ClearUser() {
	User = UserSession{
		LoggedIn: false,
	}
}
