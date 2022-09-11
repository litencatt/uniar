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
	"log"
	"os"

	"github.com/litencatt/uniar/repository"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/xo/dburl"
)

// memberCmd represents the member command
var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "List member",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		dsn := os.Getenv("UNIAR_DSN")
		db, err := dburl.Open(dsn)
		if err != nil {
			log.Print(err)
		}

		queries := repository.New(db)
		m, err := queries.GetMembers(ctx)
		if err != nil {
			log.Print(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"group", "member_id", "name", "phase", "graduated"})

		for _, v := range m {
			g := []string{
				v.Group,
				fmt.Sprintf("%d", v.MemberID),
				v.Name,
				fmt.Sprintf("%d", v.Phase),
				fmt.Sprintf("%t", v.Graduated),
			}
			table.Append(g)
		}
		table.Render()
	},
}

func init() {
	listCmd.AddCommand(memberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// memberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// memberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
