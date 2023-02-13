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

func (x *LoginProducer) AuthHandler(c *gin.Context) {
	ctx := context.Background()
	fmt.Println("AuthHandler()")
	// ログイン済みならトップにリダイレクト
	if User.LoggedIn {
		fmt.Println("リダイレクted")
		c.Redirect(http.StatusMovedPermanently, "/auth/members")
		c.Abort()
	}

	// 未ログイン時はOAuth認証後のリダイレクト時のパラメータより
	// ログイン処理を実行
	val := c.MustGet("user")
	if user, ok := val.(goauth.Userinfo); ok {
		User.Id = user.Id
		User.EMail = user.Email
		User.LoggedIn = true

		p, err := x.ProducerService.FindProducer(ctx, User.Id)
		switch {
		case err == sql.ErrNoRows:
			fmt.Printf("FindProducer record not found. User.Id = %v\n", User.Id)
		case err != nil:
			fmt.Println("FindProducer error")
			c.AbortWithError(http.StatusInternalServerError, err)
		default:
			fmt.Printf("FindProducer record found. User.Id = %v\n", User.Id)
		}

		if p.IdentityId == "" {
			if err := x.ProducerService.RegistProducer(ctx, User.Id, ""); err != nil {
				fmt.Printf("UpdateProducer error. User.Id = %v\n", User.Id)
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}

		session := sessions.Default(c)
		session.Set("ProducerId", &User.Id)
		// session.Set("Email", &User.EMail)
		session.Save()
		// fmt.Println("Saved session")
	} else {
		fmt.Println("Not Authorized")
	}

	c.Redirect(http.StatusMovedPermanently, "/auth/members")
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	pid := session.Get("ProducerId")
	session.Clear()
	session.Save()
	ClearUser()

	fmt.Printf("Logouted. User.Id = %v\n", pid)
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

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthCheck()")
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

func isLoggedIn() bool {
	return User.LoggedIn
}

func ClearUser() {
	User = UserSession{}
	User.LoggedIn = false
}
