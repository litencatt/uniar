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
	"os"
	"strconv"

	"github.com/Songmu/prompter"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

var setupSceneCmd = &cobra.Command{
	Use: "scene",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		q := repository.New()

		if err := setupScene(ctx, db, q); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func setupScene(ctx context.Context, db *sql.DB, q *repository.Queries) error {
	ps, err := q.GetProducerScenes(ctx, db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("== 所持シーンカードセットアップ ==\n0:未所持\n1:所持\nデフォルト値:0(未所持)\n\n")
	for _, s := range ps {
		ssrp := ""
		if s.SsrPlus == 1 {
			ssrp = "(SSR+)"
		}
		h := "未所持"
		if s.Have.Int64 == 1 {
			h = "1(所持)"
		}
		str := fmt.Sprintf("%s %s %s%s\n(現在: %s)", s.Photograph, s.Color, s.Member, ssrp, h)
		have := (&prompter.Prompter{
			Message: str,
			Choices: []string{"0", "1"},
			Default: fmt.Sprintf("%d", s.Have.Int64),
		}).Prompt()
		hi, _ := strconv.Atoi(have)
		if err := q.UpdateProducerScene(ctx, db, repository.UpdateProducerSceneParams{
			Have: sql.NullInt64{Int64: int64(hi), Valid: true},
			ID:   s.ID,
		}); err != nil {
			return err
		}
	}
}

func init() {
	setupCmd.AddCommand(setupSceneCmd)
}
