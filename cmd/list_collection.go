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
)

// collectionCmd represents the collection command
var listCollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "List collection",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		db, err := repository.NewConnection()
		if err != nil {
			log.Print(err)
		}

		q := repository.New()
		m, err := q.GetCollections(ctx, db)
		if err != nil {
			log.Print(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"photograph", "member", "color", "expected_value", "ssr+"})

		for _, v := range m {
			g := []string{
				v.Photograph,
				v.Member,
				v.Color,
				v.ExpectedValue.String,
				fmt.Sprintf("%t", v.SsrPlus),
			}
			table.Append(g)
		}
		table.Render()
	},
}

func init() {
	listCmd.AddCommand(listCollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
