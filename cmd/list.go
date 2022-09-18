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
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

var (
	scoreReg = regexp.MustCompile(`Score$`)
)

func render(data any) {
	// // SetHeaderのためにフィールド名取得
	var fields []string
	var scoreIndex []int
	d := reflect.TypeOf(data).Elem()
	for i := 0; i < d.NumField(); i++ {
		field := d.Field(i)
		// *SCOREはskip
		if scoreReg.MatchString(field.Name) {
			scoreIndex = append(scoreIndex, i)
			continue
		}
		fields = append(fields, field.Name)
	}
	//fmt.Printf("%v\n", fields)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fields)

	// 各Rowの値を設定
	var b [][]string
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		rv := v.Index(i)
		//fmt.Printf("v\n", rv.NumField())

		var r []string
		for i := 0; i < rv.NumField(); i++ {
			if include(scoreIndex, i) {
				continue
			}
			field := rv.Field(i)
			r = append(r, fmt.Sprintf("%v", field))
		}
		b = append(b, r)
	}
	table.AppendBulk(b)
	table.Render()
}

func include(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}
