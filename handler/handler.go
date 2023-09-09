package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	goauth "google.golang.org/api/oauth2/v2"
)

var User UserSession

type UserSession struct {
	Id       string
	EMail    string
	LoggedIn bool
}

type LoginProducer struct {
	ProducerService ProducerService
	MemberService   MemberService
}

func SetupSession(secret []byte, sessionName string) gin.HandlerFunc {
	store := cookie.NewStore(secret)
	return sessions.Sessions(sessionName, store)
}

func LoginHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/auth")
}

func RootHandler(c *gin.Context) {
	//session := sessions.Default(c)
	//es := session.Get("uniar_oauth_session")
	//fmt.Printf("%+v\n", es.(goauth.Userinfo))
	c.Redirect(http.StatusFound, "/login")
	// c.HTML(http.StatusOK, "top/index.go.tmpl", gin.H{
	// 	"title":    "Main Index",
	// 	"LoggedIn": User.LoggedIn,
	// 	"EMail":    User.EMail,
	// })
}

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("LoginAuth()")
		user := goauth.Userinfo{
			Id:    "100",
			Email: "foo@example.com",
		}
		c.Set("user", user)
	}
}

// AuthCheck is middleware for checking login status.
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isLoggedIn() {
			fmt.Println("AuthCheck() is logged in")
			val := c.MustGet("user")
			if user, ok := val.(goauth.Userinfo); ok {
				User.Id = user.Id
				User.EMail = user.Email
			}
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
		User.Id = user.Id
		User.EMail = user.Email
		User.LoggedIn = true

		p, err := x.ProducerService.FindProducer(ctx, User.Id)
		switch {
		case err == sql.ErrNoRows:
			fmt.Printf("FindProducer record not found. User.Id = %v\n", User.Id)
		case err != nil:
			fmt.Println("FindProducer error")
			fmt.Printf("%v\n", err)
			c.AbortWithError(http.StatusInternalServerError, err)
		default:
			fmt.Printf("FindProducer record found. User.Id = %v\n", User.Id)
		}

		if p.IdentityId == "" {
			if err := x.ProducerService.RegistProducer(ctx, User.Id); err != nil {
				fmt.Printf("UpdateProducer error. User.Id = %v\n", User.Id)
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}

		session.Set("uniar_oauth_user", User)
		session.Save()
		ts := session.Get("uniar_oauth_user")
		fmt.Printf("saved session: %+v\n", ts)
	} else {
		fmt.Println("Not Authorized")
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	c.Redirect(http.StatusMovedPermanently, "/auth/members")
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	// session key name of zalando/gin-oauth2
	// https://github.com/zalando/gin-oauth2/blob/master/google/google.go
	session.Set("ginoauth_google_session", "")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Save()
	ClearUser()

	c.Redirect(http.StatusMovedPermanently, "/")
}

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(User)
		if isLoggedIn() {
			val := c.MustGet("user")
			if user, ok := val.(goauth.Userinfo); ok {
				User.Id = user.Id
				User.EMail = user.Email
			}
		} else {
			ClearUser()
		}
	}
}

func isLoggedIn() bool {
	return User.LoggedIn
}

func ClearUser() {
	User = UserSession{}
	User.LoggedIn = false
}
