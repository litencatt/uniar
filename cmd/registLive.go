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

// liveCmd represents the live command
var registLiveCmd = &cobra.Command{
	Use:   "live",
	Short: "Regist a live to database",
	Run: func(cmd *cobra.Command, args []string) {
		g := (&prompter.Prompter{
			Choices: []string{"1", "2"},
			Default: "1",
			Message: "Select Group(1:櫻坂46 2:日向坂46)",
		}).Prompt()
		groupId, _ := strconv.Atoi(g)

		liveName := (&prompter.Prompter{
			Message: "Live name",
		}).Prompt()

		ctx := context.Background()
		db, err := repository.NewConnection()
		if err != nil {
			fmt.Println(err)
			return
		}

		q := repository.New()
		if err := q.RegistLive(ctx, db, liveName); err != nil {
			fmt.Println(err)
			fmt.Println("please setup first.\n$ uniar setup")
			return
		}

		gn := "櫻坂46"
		if groupId == 2 {
			gn = "日向坂46"
		}
		fmt.Printf("success registration(GroupName:%s LiveName:%s)\n", gn, liveName)
	},
}

func init() {
	registCmd.AddCommand(registLiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// liveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// liveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
