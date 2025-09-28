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
	IsAdmin  bool
}

type LoginProducer struct {
	ProducerService ProducerService
	MemberService   MemberService
}

func RootHandler(c *gin.Context) {
	fmt.Printf("RootHandler() request path: %s\n", c.Request.URL.Path)
	c.Redirect(http.StatusFound, "/login")
}

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("LoginCheck() start User:%+v\n", User)
		us, err := getUserSession(c)
		if err != nil {
			fmt.Println("LoginCheck() user session not found")
			return
		}
		fmt.Printf("key:uniar_session value: %+v\n", us)

		if us.LoggedIn {
			fmt.Println("LoginCheck() is logged in")
			c.Redirect(http.StatusFound, "/auth/members")
		} else {
			fmt.Println("LoginCheck() is NOT logged in")
		}
	}
}

// AuthCheck is middleware for checking login status.
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthCheck() start")
		fmt.Printf("AuthCheck() request path: %s\n", c.Request.URL.Path)
		if c.Request.URL.Path == "/auth/" {
			fmt.Println("AuthCheck() request path is /auth/")
			return
		}

		us, err := getUserSession(c)
		if err != nil {
			fmt.Println("AuthCheck() user session not found")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		fmt.Printf("key:uniar_session value: %+v\n", us)

		if us.LoggedIn {
			fmt.Println("AuthCheck() is logged in")
			c.Set("user", us)
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
	us := session.Get("uniar_session")
	fmt.Printf("AuthHandler() uniar_session: %+v\n", us)
	if us != nil && us.(UserSession).LoggedIn {
		fmt.Println("already Loggedin. redirect to /auth/members")
		c.Redirect(http.StatusMovedPermanently, "/auth/members")
	} 
	fmt.Println("AuthHandler() uniar_session not found")

	// contextに保存されたGoogle認証情報を取得
	ctxUser := c.MustGet("user")
	if goauthUser, ok := ctxUser.(goauth.Userinfo); ok {
		fmt.Printf("AuthHandler() goauthUser.Id:%v, Email:%v\n", goauthUser.Id, goauthUser.Email)
		p, err := x.ProducerService.FindProducer(ctx, goauthUser.Id)
		fmt.Printf("AuthHandler() find producer result:%+v\n", p)
		switch {
		case err == sql.ErrNoRows:
			fmt.Printf("AuthHandler() producer not found. goauthUser.Id = %v\n", goauthUser.Id)
			if err := x.ProducerService.RegistProducer(ctx, goauthUser.Id); err != nil {
				fmt.Printf("regist producer error. goauthUser.Id = %v\n", goauthUser.Id)
				if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
				fmt.Printf("AbortWithError failed: %v\n", err)
			}
			}
			// 新規登録後に再度プロデューサー情報を取得
			p, err = x.ProducerService.FindProducer(ctx, goauthUser.Id)
			if err != nil {
				fmt.Printf("AuthHandler() find producer after regist error: %v\n", err)
				if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
					fmt.Printf("AbortWithError failed: %v\n", err)
				}
				return
			}
			fmt.Printf("AuthHandler() producer found after regist: %+v\n", p)
		case err != nil:
			fmt.Printf("find producer error: goauthUser.Id = %v err = %v\n", goauthUser.Id, err)
			if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
				fmt.Printf("AbortWithError failed: %v\n", err)
			}
		default:
			fmt.Printf("AuthHandler() producer record found. goauthUser.Id = %v\n", goauthUser.Id)
		}
		userSession := UserSession{
			IdentityId: goauthUser.Id,
			EMail:      goauthUser.Email,
			LoggedIn:   true,
			ProducerId: p.ID,
			IsAdmin:    p.IsAdmin,
		}
		fmt.Printf("AuthHandler() save uniar_session value:%+v\n", userSession)
		session.Set("uniar_session", &userSession)
		if err := session.Save(); err != nil {
			fmt.Printf("AuthHandler() session save error: %+v\n", err)
			if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
				fmt.Printf("AbortWithError failed: %v\n", err)
			}
		}

		us := session.Get("uniar_session")
		fmt.Printf("AuthHandler() get uniar_session value:%+v\n", us)
	} else {
		fmt.Println("AuthHandler() Not Authorized")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/auth/members")
}

func LogoutHandler(c *gin.Context) {
	fmt.Println("LogoutHandler() start")

	c.Set("user", nil)
	session := sessions.Default(c)
	// session key name of zalando/gin-oauth2
	// https://github.com/zalando/gin-oauth2/blob/master/google/google.go
	session.Delete("ginoauth_google_session")
	session.Delete("uniar_session")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	if err := session.Save(); err != nil {
		fmt.Printf("LogoutHandler() session save error: %v\n", err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

// AdminCheck is middleware for checking admin privileges.
func AdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AdminCheck() start")
		us, err := getUserSession(c)
		if err != nil {
			fmt.Println("AdminCheck() user session not found")
			c.HTML(http.StatusForbidden, "error/403.go.tmpl", gin.H{
				"message": "管理者権限が必要です",
			})
			c.Abort()
			return
		}

		if !us.IsAdmin {
			fmt.Printf("AdminCheck() user is not admin: %+v\n", us)
			c.HTML(http.StatusForbidden, "error/403.go.tmpl", gin.H{
				"message": "管理者権限が必要です",
			})
			c.Abort()
			return
		}

		fmt.Println("AdminCheck() admin access granted")
		c.Next()
	}
}

func getUserSession(c *gin.Context) (*UserSession, error) {
	session := sessions.Default(c)
	us := session.Get("uniar_session")
	fmt.Printf("getUserSession() key:uniar_session value: %+v\n", us)

	if us == nil {
		return nil, fmt.Errorf("getUserSession() user session not found")
	}
	return us.(*UserSession), nil
}