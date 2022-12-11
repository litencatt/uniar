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
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var registSceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "Regist a scene to database",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			return err
		}

		gid := (&prompter.Prompter{
			Message: "Select GroupID(1:櫻坂46 2:日向坂46)",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		pt := (&prompter.Prompter{
			Message: "Select Type",
			Choices: []string{"楽曲", "限定", "Precious"},
			Default: "楽曲",
		}).Prompt()

		groupId, _ := strconv.Atoi(gid)
		q := repository.New()
		pl, err := q.GetPhotographList(ctx, db, repository.GetPhotographListParams{
			GroupID:   int64(groupId),
			PhotoType: pt,
		})
		if err != nil {
			log.Print(err)
		}

		var pList []string
		for _, v := range pl {
			pList = append(pList, fmt.Sprintf("%d: %s", v.ID, v.Name))
		}
		pListStr := strings.Join(pList, "\n")
		pid := (&prompter.Prompter{
			Message: fmt.Sprintf("Select Photograph ID\n%s\n", pListStr),
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		ml, err := q.GetMemberList(ctx, db, int64(groupId))
		if err != nil {
			return err
		}
		var mList []string
		for _, v := range ml {
			mList = append(mList, fmt.Sprintf("%d: %s", v.ID, v.Name))
		}
		mListStr := strings.Join(mList, "\n")
		mid := (&prompter.Prompter{
			Message: fmt.Sprintf("Select Member ID\n%s\n", mListStr),
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		col := (&prompter.Prompter{
			Message: "Color Type(1:Red, 2:Blue, 3:Green, 4:Yellow, 5:Purple, 6:All",
			Choices: []string{"1", "2", "3", "4", "5", "6"},
		}).Prompt()

		exp := (&prompter.Prompter{
			Message: "Expected Value",
		}).Prompt()

		vo := (&prompter.Prompter{
			Message: "Vocal Max",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		da := (&prompter.Prompter{
			Message: "Dance Max",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		pe := (&prompter.Prompter{
			Message: "Peformance Max",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		sp := (&prompter.Prompter{
			Message: "Is a SSR+?",
			Choices: []string{"0", "1"},
			Default: "0",
		}).Prompt()

		photoId, _ := strconv.Atoi(pid)
		memberId, _ := strconv.Atoi(mid)
		colId, _ := strconv.Atoi(col)
		voMax, _ := strconv.Atoi(vo)
		daMax, _ := strconv.Atoi(da)
		peMax, _ := strconv.Atoi(pe)
		ssrPlus, _ := strconv.Atoi(sp)

		expVal := sql.NullString{String: "", Valid: false}
		if exp != "" {
			expVal = sql.NullString{String: exp, Valid: true}
		}

		if err := q.RegistScene(ctx, db, repository.RegistSceneParams{
			PhotographID:   int64(photoId),
			MemberID:       int64(memberId),
			ColorTypeID:    int64(colId),
			ExpectedValue:  expVal,
			VocalMax:       int64(voMax),
			DanceMax:       int64(daMax),
			PerformanceMax: int64(peMax),
			SsrPlus:        int64(ssrPlus),
		}); err != nil {
			return err
		}

		fmt.Println("success registration")
		return nil
	},
}

func init() {
	registCmd.AddCommand(registSceneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sceneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sceneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
