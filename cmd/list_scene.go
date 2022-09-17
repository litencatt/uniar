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

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var listSceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "List scene",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := cmd.Flags().GetString("color")
		if err != nil {
			log.Fatal(err)
		}

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
				all35 := float32(v.Total) * 1.35
				all40 := float32(v.Total) * 1.40
				voda50 := float32(v.VocalMax+v.DanceMax)*1.50 + float32(v.PeformanceMax)
				dape50 := float32(v.DanceMax+v.PeformanceMax)*1.50 + float32(v.VocalMax)
				vope50 := float32(v.VocalMax+v.PeformanceMax)*1.50 + float32(v.DanceMax)
				vo85 := float32(v.VocalMax)*1.85 + float32(v.DanceMax) + float32(v.PeformanceMax)
				da85 := float32(v.DanceMax)*1.85 + float32(v.VocalMax) + float32(v.PeformanceMax)
				pe85 := float32(v.PeformanceMax)*1.85 + float32(v.VocalMax) + float32(v.DanceMax)

				scene := entity.Scene{
					Photograph:    v.Photograph,
					Member:        v.Member,
					Color:         v.Color,
					Total:         v.Total,
					All35:         int32(all35),
					All40:         int32(all40),
					VoDa50:        int32(voda50),
					DaPe50:        int32(dape50),
					VoPe50:        int32(vope50),
					Vo85:          int32(vo85),
					Da85:          int32(da85),
					Pe85:          int32(pe85),
					VocalMax:      v.VocalMax,
					DanceMax:      v.DanceMax,
					PeformanceMax: v.PeformanceMax,
					ExpectedValue: v.ExpectedValue.String,
				}
				scenes = append(scenes, scene)
			}
		} else {
			s, err := q.GetScenesWithColor(ctx, db, c)
			if err != nil {
				log.Print(err)
			}
			for _, v := range s {
				all35 := float32(v.Total) * 1.35
				all40 := float32(v.Total) * 1.40
				voda50 := float32(v.VocalMax+v.DanceMax)*1.50 + float32(v.PeformanceMax)
				dape50 := float32(v.DanceMax+v.PeformanceMax)*1.50 + float32(v.VocalMax)
				vope50 := float32(v.VocalMax+v.PeformanceMax)*1.50 + float32(v.DanceMax)
				vo85 := float32(v.VocalMax)*1.85 + float32(v.DanceMax) + float32(v.PeformanceMax)
				da85 := float32(v.DanceMax)*1.85 + float32(v.VocalMax) + float32(v.PeformanceMax)
				pe85 := float32(v.PeformanceMax)*1.85 + float32(v.VocalMax) + float32(v.DanceMax)

				scene := entity.Scene{
					Photograph:    v.Photograph,
					Member:        v.Member,
					Color:         v.Color,
					Total:         v.Total,
					All35:         int32(all35),
					All40:         int32(all40),
					VoDa50:        int32(voda50),
					DaPe50:        int32(dape50),
					VoPe50:        int32(vope50),
					Vo85:          int32(vo85),
					Da85:          int32(da85),
					Pe85:          int32(pe85),
					VocalMax:      v.VocalMax,
					DanceMax:      v.DanceMax,
					PeformanceMax: v.PeformanceMax,
					ExpectedValue: v.ExpectedValue.String,
				}
				scenes = append(scenes, scene)
			}
		}

		render(scenes)
	},
}

func init() {
	listCmd.AddCommand(listSceneCmd)
	listSceneCmd.Flags().StringP("color", "c", "", "Color filter")
}
