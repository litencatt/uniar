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

		var scenes any
		q := repository.New()
		if c == "" {
			scenes, err = q.GetScenes(ctx, db)
		} else {
			scenes, err = q.GetScenesWithColor(ctx, db, c)
		}
		if err != nil {
			log.Print(err)
		}

		render(scenes)
	},
}

func init() {
	listCmd.AddCommand(listSceneCmd)
	listSceneCmd.Flags().StringP("color", "c", "", "Color filter")
}
