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
	"regexp"
	"strconv"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

// musicCmd represents the music command
var registMusicCmd = &cobra.Command{
	Use:   "music",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		db, err := repository.NewConnection()
		if err != nil {
			log.Print(err)
		}
		q := repository.New()
		ll, err := q.GetLiveList(ctx, db)
		if err != nil {
			log.Print(err)
		}
		var liveList []string
		for _, v := range ll {
			liveList = append(liveList, fmt.Sprintf("%d: %s", v.ID, v.Name))
		}
		liveListStr := strings.Join(liveList, "\n")

		lid := (&prompter.Prompter{
			Message: fmt.Sprintf("Select Live ID\n%s\n", liveListStr),
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		name := (&prompter.Prompter{
			Message: "Music Name",
		}).Prompt()

		nd := (&prompter.Prompter{
			Message: "Normal difficulty",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		pd := (&prompter.Prompter{
			Message: "Pro difficulty",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		md := (&prompter.Prompter{
			Message: "Master difficulty",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		len := (&prompter.Prompter{
			Message: "Music length(sec.)",
			Regexp:  regexp.MustCompile(`\d`),
		}).Prompt()

		col := (&prompter.Prompter{
			Message: "Color Type(1:Red, 2:Blue, 3:Green, 4:Yellow, 5:Purple, 6:All",
			Choices: []string{"1", "2", "3", "4", "5", "6"},
		}).Prompt()

		ndi, _ := strconv.Atoi(nd)
		pdi, _ := strconv.Atoi(pd)
		mdi, _ := strconv.Atoi(md)
		sec, _ := strconv.Atoi(len)
		colId, _ := strconv.Atoi(col)
		liveId, _ := strconv.Atoi(lid)
		if err := q.RegistMusic(ctx, db, repository.RegistMusicParams{
			Name:        name,
			Normal:      int64(ndi),
			Pro:         int64(pdi),
			Master:      int64(mdi),
			Length:      int64(sec),
			ColorTypeID: int64(colId),
			LiveID:      int64(liveId),
		}); err != nil {
			log.Print(err)
		}

		fmt.Printf("success registration(MusicName:%s)\n", name)
	},
}

func init() {
	registCmd.AddCommand(registMusicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// musicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// musicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
