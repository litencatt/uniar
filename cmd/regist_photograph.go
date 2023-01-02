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
	"fmt"
	"strconv"

	"github.com/Songmu/prompter"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

// photographCmd represents the photograph command
var registPhotographCmd = &cobra.Command{
	Use:     "photo",
	Short:   "Regist a photograph to database",
	Aliases: []string{"p"},
	RunE: func(cmd *cobra.Command, args []string) error {
		g := (&prompter.Prompter{
			Choices: []string{"1", "2"},
			Default: "1",
			Message: "Select Group(1:櫻坂46 2:日向坂46)",
		}).Prompt()
		groupId, _ := strconv.Atoi(g)

		pt := (&prompter.Prompter{
			Message: "Select Type",
			Choices: []string{"楽曲", "限定", "Precious"},
			Default: "楽曲",
		}).Prompt()

		photoName := (&prompter.Prompter{
			Message: "Photograph Name",
		}).Prompt()

		ctx := context.Background()
		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			fmt.Println(err)
		}

		q := repository.New()
		if err := q.RegistPhotograph(ctx, db, repository.RegistPhotographParams{
			Name:      photoName,
			GroupID:   int64(groupId),
			PhotoType: pt,
		}); err != nil {
			return err
		}

		fmt.Printf("success registration(Photograph Name: %s)\n", photoName)
		return nil
	},
}

func init() {
	registCmd.AddCommand(registPhotographCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// photographCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// photographCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
