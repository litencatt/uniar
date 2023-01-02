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
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

var listSceneCmd = &cobra.Command{
	Use:     "scene",
	Short:   "Show scene card list",
	Aliases: []string{"s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// User inputs
		c, _ := cmd.Flags().GetString("color")
		c = getColorName(c)
		m, _ := cmd.Flags().GetString("member")
		p, _ := cmd.Flags().GetString("photograph")
		s, _ := cmd.Flags().GetString("sort")
		h, _ := cmd.Flags().GetBool("have")
		n, _ := cmd.Flags().GetBool("not-have")
		d, _ := cmd.Flags().GetBool("detail")
		f, _ := cmd.Flags().GetBool("full-name")

		ctx := context.Background()
		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			return err
		}
		q := repository.New()

		var scenes []entity.Scene
		color := c
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
		ss, err := q.GetScenesWithColor(ctx, db, repository.GetScenesWithColorParams{
			Name:   color,
			Name_2: member,
			Name_3: photo,
		})
		if err != nil {
			return err
		}
		for _, s := range ss {
			// Show only scene you have
			if h && s.Have == 0 {
				continue
			}
			// Show only scene you not have
			if n && s.Have != 0 {
				continue
			}

			var e float64
			if s.ExpectedValue.Valid {
				e, _ = strconv.ParseFloat(s.ExpectedValue.String, 32)
			}
			p := s.Photograph
			if !f && s.Abbreviation != "" {
				p = s.Abbreviation
			}
			scene := entity.Scene{
				Photograph: p,
				Member:     s.Member,
				Color:      s.Color,
				Total:      s.Total,
				Vo:         s.VocalMax,
				Da:         s.DanceMax,
				Pe:         s.PerformanceMax,
				Expect:     float32(e),
				SsrPlus:    s.SsrPlus == 1,
			}
			scene.CalcTotal(s.Bonds, s.Discography)
			scenes = append(scenes, scene)
		}

		sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
		for i, _ := range scenes {
			scenes[i].All35 = int64(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoDa50Score > scenes[j].VoDa50Score })
		for i, _ := range scenes {
			scenes[i].VoDa50 = int64(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].DaPe50Score > scenes[j].DaPe50Score })
		for i, _ := range scenes {
			scenes[i].DaPe50 = int64(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoPe50Score > scenes[j].VoPe50Score })
		for i, _ := range scenes {
			scenes[i].VoPe50 = int64(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Vo85Score > scenes[j].Vo85Score })
		for i, _ := range scenes {
			scenes[i].Vo85 = int64(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Da85Score > scenes[j].Da85Score })
		for i, _ := range scenes {
			scenes[i].Da85 = int64(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Pe85Score > scenes[j].Pe85Score })
		for i, _ := range scenes {
			scenes[i].Pe85 = int64(i + 1)
		}

		switch s {
		case "all35":
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
		case "voda50":
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoDa50Score > scenes[j].VoDa50Score })
		case "dape50":
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].DaPe50Score > scenes[j].DaPe50Score })
		case "vope50":
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoPe50Score > scenes[j].VoPe50Score })
		case "vo85":
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].Vo85Score > scenes[j].Vo85Score })
		case "da85":
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].Da85Score > scenes[j].Da85Score })
		case "pe85":
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].Pe85Score > scenes[j].Pe85Score })
		default:
			sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
		}

		ignoreColumnsStr, _ := cmd.Flags().GetString("ignore-columns")
		if !d && ignoreColumnsStr != "" {
			ignoreColumnsStr += ",Vo,Da,Pe"
		}
		if !d && ignoreColumnsStr == "" {
			ignoreColumnsStr = "Vo,Da,Pe"
		}
		ic := strings.Split(ignoreColumnsStr, ",")

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
