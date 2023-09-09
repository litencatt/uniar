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

func RootHandler(c *gin.Context) {
	if isLoggedIn() {
		c.Redirect(http.StatusFound, "/auth/members")
	}
	c.Redirect(http.StatusFound, "/login")
}

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isLoggedIn() {
			fmt.Println("LoginCheck() is logged in")
			c.Redirect(http.StatusFound, "/auth")
		} else {
			fmt.Println("LoginCheck() is NOT logged in")
		}
	}
}

// AuthCheck is middleware for checking login status.
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		User.Id = user.Id
		User.EMail = user.Email
		User.LoggedIn = true

		p, err := x.ProducerService.FindProducer(ctx, User.Id)
		switch {
		case err == sql.ErrNoRows:
			fmt.Printf("producer not found. User.Id = %v\n", User.Id)
		case err != nil:
			fmt.Printf("find producer error: User.Id = %v err = %v\n", User.Id, err)
			c.AbortWithError(http.StatusInternalServerError, err)
		default:
			fmt.Printf("producer record found. User.Id = %v\n", User.Id)
		}

		if p.IdentityId == "" {
			if err := x.ProducerService.RegistProducer(ctx, User.Id); err != nil {
				fmt.Printf("update producer error. User.Id = %v\n", User.Id)
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
