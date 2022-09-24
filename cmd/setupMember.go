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
	"regexp"
	"strconv"

	"github.com/Songmu/prompter"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

var setupMemberCmd = &cobra.Command{
	Use: "member",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		q := repository.New()

		if err := setupMember(ctx, db, q); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func setupMember(ctx context.Context, db *sql.DB, q *repository.Queries) error {
	fmt.Printf("== メンバーステータスセットアップ ==\n")
	pm, _ := q.GetProducerMember(ctx, db)
	for _, m := range pm {
		fmt.Println(m.Name)
		pmb := (&prompter.Prompter{
			Message: fmt.Sprintf("絆ランク (現在値:%d)", m.BondLevelCurent),
			Regexp:  regexp.MustCompile(`^\d*$`),
			Default: fmt.Sprintf("%d", m.BondLevelCurent),
		}).Prompt()
		pmbi, _ := strconv.Atoi(pmb)

		ddt := (&prompter.Prompter{
			Message: fmt.Sprintf("ディスコグラフィ (現在値:%d)", m.DiscographyDiscTotal),
			Regexp:  regexp.MustCompile(`^\d*$`),
			Default: fmt.Sprintf("%d", m.DiscographyDiscTotal),
		}).Prompt()
		ddti, _ := strconv.Atoi(ddt)

		if err := q.UpdateProducerMember(ctx, db, repository.UpdateProducerMemberParams{
			ID:                   m.ID,
			BondLevelCurent:      int64(pmbi),
			DiscographyDiscTotal: int64(ddti),
		}); err != nil {
			return err
		}
		fmt.Println()
	}
	return nil
}

func init() {
	setupCmd.AddCommand(setupMemberCmd)
}
