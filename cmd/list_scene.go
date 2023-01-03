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
	"os"
	"strings"

	"github.com/litencatt/uniar/repository"
	"github.com/litencatt/uniar/service"
	"github.com/spf13/cobra"
)

var listSceneCmd = &cobra.Command{
	Use:     "scene",
	Short:   "Show scene card list",
	Aliases: []string{"s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _ := cmd.Flags().GetString("color")
		m, _ := cmd.Flags().GetString("member")
		p, _ := cmd.Flags().GetString("photograph")
		s, _ := cmd.Flags().GetString("sort")
		h, _ := cmd.Flags().GetBool("have")
		n, _ := cmd.Flags().GetBool("not-have")
		d, _ := cmd.Flags().GetBool("detail")
		f, _ := cmd.Flags().GetBool("full-name")
		i, _ := cmd.Flags().GetString("ignore-columns")

		// ユーザー入力の整形
		color := getColorName(c)
		if c == "" {
			color = "%"
		}

		member := m
		if m == "" {
			member = "%"
		} else {
			member = "%" + m + "%"
		}

		photo := p
		if p == "" {
			photo = "%"
		} else {
			photo = "%" + p + "%"
		}

		// 指定非表示カラム
		if !d && i != "" {
			i += ",Vo,Da,Pe"
		}
		if !d && i == "" {
			i = "Vo,Da,Pe"
		}
		ic := strings.Split(i, ",")

		ctx := context.Background()

		req := service.ListSceneRequest{
			Color:      color,
			Member:     member,
			Photograph: photo,
			Sort:       s,
			Have:       h,
			NotHave:    n,
			Detail:     d,
			FullName:   f,
		}

		db, err := repository.NewConnection(GetDbPath())
		if err != nil {
			return err
		}
		q := repository.New()

		svc := service.ListScene{
			DB:      db,
			Querier: q,
		}
		scenes, err := svc.ListScene(ctx, &req)
		if err != nil {
			return err
		}

		render(os.Stdout, scenes, ic)
		return nil
	},
}

func getColorName(c string) string {
	switch c {
	case "r", "red":
		return "Red"
	case "b", "blue":
		return "Blue"
	case "g", "green":
		return "Green"
	case "y", "yellow":
		return "Yellow"
	case "p", "purple":
		return "Purple"
	default:
		return c
	}
}

func init() {
	listCmd.AddCommand(listSceneCmd)
	listSceneCmd.Flags().BoolP("have", "", false, "Show only scenes you have")
	listSceneCmd.Flags().BoolP("not-have", "n", false, "Show only scenes you NOT have")
	listSceneCmd.Flags().BoolP("detail", "d", false, "Show detail")
	listSceneCmd.Flags().BoolP("full-name", "f", false, "Show pohtograph full name")
	listSceneCmd.Flags().StringP("color", "c", "", "Color filter(e.g. -c Red or -c r)")
	listSceneCmd.Flags().StringP("member", "m", "", "Member filter(e.g. -m 加藤史帆)")
	listSceneCmd.Flags().StringP("photograph", "p", "", "Photograph filter(e.g. -p JOYFULLOVE)")
	listSceneCmd.Flags().StringP("sort", "s", "", "Sort target rank.(all35, voda50, ...)")
	listSceneCmd.Flags().StringP("ignore-columns", "i", "", "Ignore columns to display(VoDa50,DaPe50,...)")
}
