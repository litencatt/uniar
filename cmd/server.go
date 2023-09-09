/*
Copyright © 2022 Kosuke Nakamura <ncl0709@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/handler"
	"github.com/litencatt/uniar/repository"
	"github.com/litencatt/uniar/service"
	"github.com/spf13/cobra"
	"github.com/zalando/gin-oauth2/google"
)

const (
	// session name for CookieStore
	sessionName = "uniar_oauth_session"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Server",
	Aliases: []string{"s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := run(context.Background()); err != nil {
			log.Printf("failed to terminate server: %v", err)
			os.Exit(1)
		}
		return nil
	},
}

func run(ctx context.Context) error {
	db, err := repository.NewConnection(GetDbPath())
	if err != nil {
		return err
	}
	q := repository.New()

	ls := &handler.ListScene{
		SceneService:      &service.Scene{DB: db, Querier: q},
		MemberService:     &service.Member{DB: db, Querier: q},
		PhotographService: &service.Photgraph{DB: db, Querier: q},
	}

	rs := &handler.RegistScene{
		SceneService:         &service.Scene{DB: db, Querier: q},
		ProducerSceneService: &service.ProducerScene{DB: db, Querier: q},
		MemberService:        &service.Member{DB: db, Querier: q},
		PhotographService:    &service.Photgraph{DB: db, Querier: q},
	}

	lm := &handler.ListMember{
		MemberService: &service.Member{DB: db, Querier: q},
	}
	ah := &handler.LoginProducer{
		ProducerService: &service.Producer{DB: db, Querier: q},
	}

	redirectURL := "http://127.0.0.1:8090/auth"
	credFile := "./cred.json"
	scopes := []string{"https://www.googleapis.com/auth/userinfo.email"}
	secret := []byte("secret")

	r := gin.Default()
	google.Setup(redirectURL, credFile, scopes, secret)
	r.Use(google.Session(sessionName))

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", handler.RootHandler)

	login := r.Group("/login")
	{
		login.Use(handler.LoginCheck())
		login.GET("/", google.LoginHandler)
	}

	// /auth 以下は認証が必要
	private := r.Group("/auth")
	{
		private.Use(google.Auth())
		private.Use(handler.AuthCheck())
	
		private.GET("/", ah.AuthHandler)
		private.GET("/logout", handler.LogoutHandler)
	
		private.GET("/scenes", ls.ListScene)
	
		private.GET("/regist/:group_id", rs.GetRegist)
		private.POST("/regist/:group_id", rs.PostRegist)
	
		private.GET("/members", lm.ListMember)
		private.POST("/members", lm.UpdateMember)
	}

	r.Run(":8090")
	return nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
