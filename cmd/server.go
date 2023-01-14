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
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	db, err := repository.NewConnection(GetDbPath())
	if err != nil {
		return err
	}
	q := repository.New()

	r.Static("/assets", "./assets")

	r.GET("/", handler.Top)

	ls := &handler.ListScene{
		SceneService:      &service.Scene{DB: db, Querier: q},
		MemberService:     &service.Member{DB: db, Querier: q},
		PhotographService: &service.Photgraph{DB: db, Querier: q},
	}
	r.GET("/scenes", ls.ListScene)

	rs := &handler.RegistScene{
		SceneService:      &service.Scene{DB: db, Querier: q},
		MemberService:     &service.Member{DB: db, Querier: q},
		PhotographService: &service.Photgraph{DB: db, Querier: q},
	}
	r.GET("/regist", rs.GetRegist)
	r.POST("/regist", rs.PostRegist)

	r.Run(":8090")
	return nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
