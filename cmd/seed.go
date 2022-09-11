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

	"github.com/litencatt/unisonair/repository"
	"github.com/spf13/cobra"
	"github.com/xo/dburl"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		dsn := os.Getenv("UNIAR_DSN")
		db, err := dburl.Open(dsn)
		if err != nil {
			log.Print(err)
		}

		queries := repository.New(db)
		if err := queries.SeedGroups(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedCenterSkills(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedColorTypes(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedLives(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedMembers(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedMusic(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedPhotograph(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedScenes(ctx); err != nil {
			log.Print(err)
		}
		if err := queries.SeedSkills(ctx); err != nil {
			log.Print(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
