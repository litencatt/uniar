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

var setupOfficeCmd = &cobra.Command{
	Use:     "office",
	Aliases: []string{"o"},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			return err
		}
		q := repository.New()

		if err := setupOffice(ctx, db, q); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return nil
	},
}

func initProducerOffice(ctx context.Context, db *sql.DB, q *repository.Queries) error {
	producerId := 1
	if err := q.RegistProducerOffice(ctx, db, int64(producerId)); err != nil {
		return err
	}
	return nil
}

func setupOffice(ctx context.Context, db *sql.DB, q *repository.Queries) error {
	fmt.Printf("== 事務所ボーナスセットアップ ==\n")
	po, err := q.GetProducerOffice(ctx, db, 1)
	if err != nil {
		if err = initProducerOffice(ctx, db, q); err != nil {
			return err
		}
		fmt.Printf("== 事務所ボーナス初期化完了 ==\n")
		po, err = q.GetProducerOffice(ctx, db, 1)
		if err != nil {
			return err
		}
	}

	ob := (&prompter.Prompter{
		Message: fmt.Sprintf("事務所ボーナス平均値 (現在値:%d) [1-17]", po.OfficeBonus.Int64),
		Regexp:  regexp.MustCompile(`^([1-9]|1[0-7])$`),
	}).Prompt()
	obi, _ := strconv.Atoi(ob)
	uop := repository.UpdateProducerOfficeParams{
		ProducerID: 1,
		OfficeBonus: sql.NullInt64{
			Int64: int64(obi),
			Valid: true,
		},
	}
	if err := q.UpdateProducerOffice(ctx, db, uop); err != nil {
		return err
	}
	fmt.Println()
	return nil
}

func init() {
	setupCmd.AddCommand(setupOfficeCmd)
}
