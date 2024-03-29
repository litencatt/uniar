/*
Copyright © 2022 Kosuke Nakamura <litencatt@gmail.com>

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
	"strconv"

	"github.com/Songmu/prompter"
	"github.com/fatih/color"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

var setupSceneCmd = &cobra.Command{
	Use:          "scene",
	Aliases:      []string{"s"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		m, _ := cmd.Flags().GetString("member")
		p, _ := cmd.Flags().GetString("photograph")

		member := m
		if m == "" {
			member = "%"
		} else {
			member = "%" + m + "%"
		}

		photo := p
		if p == "" {
			photo = "%"
		} else {
			photo = "%" + p + "%"
		}

		ctx := context.Background()
		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			return err
		}
		q := repository.New()

		if err := setupScene(ctx, db, q, photo, member); err != nil {
			return err
		}
		return nil
	},
}

func initProducerScene(ctx context.Context, db *sql.DB, q *repository.Queries) error {
	// producer_scenesレコード作成
	ss, err := q.GetAllScenes(ctx, db)
	if err != nil {
		return err
	}
	for _, s := range ss {
		if err := q.RegistProducerScene(ctx, db, repository.RegistProducerSceneParams{
			ProducerID:   1,
			PhotographID: s.PhotographID,
			MemberID:     s.MemberID,
			SsrPlus:      s.SsrPlus,
		}); err != nil {
			return err
		}
	}

	return nil
}

func setupScene(ctx context.Context, db *sql.DB, q *repository.Queries, photo, member string) error {
	fmt.Printf("== 所持シーンカードセットアップ ==\n0:未所持\n1:所持\nデフォルト値:0(未所持)\n\n")
	ps, err := q.GetProducerScenes(ctx, db, repository.GetProducerScenesParams{
		Name:   photo,
		Name_2: member,
	})
	if err != nil {
		return err
	}

	for _, s := range ps {
		ssrp := ""
		if s.SsrPlus == 1 {
			continue
		}
		h := "未所持"
		if s.Have.Valid && s.Have.Int64 != 0 {
			h = "1(所持)"
		}

		var col string
		switch s.Color {
		case "Red":
			col = color.RedString(s.Color)
		case "Blue":
			col = color.BlueString(s.Color)
		case "Green":
			col = color.GreenString(s.Color)
		case "Yellow":
			col = color.YellowString(s.Color)
		case "Purple":
			col = color.MagentaString(s.Color)
		}

		str := fmt.Sprintf("%s %s %s%s\n(現在: %s)", s.Photograph, col, s.Member, ssrp, h)
		have := (&prompter.Prompter{
			Message: str,
			Choices: []string{"0", "1"},
			Default: fmt.Sprintf("%d", s.Have.Int64),
		}).Prompt()

		hi, _ := strconv.Atoi(have)
		if err := q.UpdateProducerScene(ctx, db, repository.UpdateProducerSceneParams{
			Have:         sql.NullInt64{Valid: true, Int64: int64(hi)},
			ProducerID:   s.ProducerID,
			PhotographID: s.PhotographID,
			MemberID:     s.MemberID,
		}); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	setupCmd.AddCommand(setupSceneCmd)
	setupSceneCmd.Flags().StringP("member", "m", "", "Member filter(e.g. -m 加藤史帆)")
	setupSceneCmd.Flags().StringP("photograph", "p", "", "Photograph filter(e.g. -p JOYFULLOVE)")
}
