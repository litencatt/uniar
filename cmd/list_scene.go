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
	"sort"
	"strconv"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var listSceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "List scene",
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := cmd.Flags().GetString("color")
		sortRank, _ := cmd.Flags().GetString("sort")

		ctx := context.Background()
		db, err := repository.NewConnection()
		if err != nil {
			log.Print(err)
		}

		var scenes []entity.Scene
		q := repository.New()
		if c == "" {
			s, err := q.GetScenes(ctx, db)
			if err != nil {
				log.Print(err)
			}
			for _, v := range s {
				var e float64
				if v.ExpectedValue.Valid {
					e, _ = strconv.ParseFloat(v.ExpectedValue.String, 32)
				}
				scene := entity.Scene{
					Photograph: v.Photograph,
					Member:     v.Member,
					Color:      v.Color,
					Total:      v.Total,
					Vo:         v.VocalMax,
					Da:         v.DanceMax,
					Pe:         v.PeformanceMax,
					Expect:     float32(e),
				}
				scene.CalcTotal(v.Bonds, v.Discography)
				scenes = append(scenes, scene)
			}
		} else {
			s, err := q.GetScenesWithColor(ctx, db, c)
			if err != nil {
				log.Print(err)
			}
			for _, v := range s {
				var e float64
				if v.ExpectedValue.Valid {
					e, _ = strconv.ParseFloat(v.ExpectedValue.String, 32)
				}
				scene := entity.Scene{
					Photograph: v.Photograph,
					Member:     v.Member,
					Color:      v.Color,
					Total:      v.Total,
					Vo:         v.VocalMax,
					Da:         v.DanceMax,
					Pe:         v.PeformanceMax,
					Expect:     float32(e),
				}
				scene.CalcTotal(v.Bonds, v.Discography)
				scenes = append(scenes, scene)
			}
		}

		sort.Slice(scenes, func(i, j int) bool { return scenes[i].All35Score > scenes[j].All35Score })
		for i, _ := range scenes {
			scenes[i].All35Rank = int32(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoDa50Score > scenes[j].VoDa50Score })
		for i, _ := range scenes {
			scenes[i].VoDa50Rank = int32(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].DaPe50Score > scenes[j].DaPe50Score })
		for i, _ := range scenes {
			scenes[i].DaPe50Rank = int32(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].VoPe50Score > scenes[j].VoPe50Score })
		for i, _ := range scenes {
			scenes[i].VoPe50Rank = int32(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Vo85Score > scenes[j].Vo85Score })
		for i, _ := range scenes {
			scenes[i].Vo85Rank = int32(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Da85Score > scenes[j].Da85Score })
		for i, _ := range scenes {
			scenes[i].Da85Rank = int32(i + 1)
		}
		sort.Slice(scenes, func(i, j int) bool { return scenes[i].Pe85Score > scenes[j].Pe85Score })
		for i, _ := range scenes {
			scenes[i].Pe85Rank = int32(i + 1)
		}

		switch sortRank {
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
		render(scenes)
	},
}

func init() {
	listCmd.AddCommand(listSceneCmd)
	listSceneCmd.Flags().StringP("color", "c", "", "Color filter")
	listSceneCmd.Flags().StringP("sort", "s", "", "Sort target rank.(all35, voda50, ...)")
}
