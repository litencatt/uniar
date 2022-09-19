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

	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

// listLiveCmd represents the listLive command
var listLiveCmd = &cobra.Command{
	Use:   "live",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		db, err := repository.NewConnection()
		if err != nil {
			fmt.Println(err)
			return
		}

		q := repository.New()
		l, err := q.GetLiveList(ctx, db)
		if err != nil {
			fmt.Println(err)
			fmt.Println("please setup first.\n$ uniar setup")
			return
		}

		render(l, []string{})
	},
}

func init() {
	listCmd.AddCommand(listLiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listLiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listLiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
