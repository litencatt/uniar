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
	"os"
	"os/user"
	"regexp"
	"strconv"

	"github.com/Songmu/prompter"
	_ "github.com/mattn/go-sqlite3"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/sqlite3"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/litencatt/uniar/repository"
	mig_sql "github.com/litencatt/uniar/sql"
	"github.com/spf13/cobra"
)

var (
	options = sqldef.Options{}
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup uniar",
	Long:  "Setup your member status and scene card collections for uniar",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		uniarPath := user.HomeDir + "/.uniar"
		dbPath := uniarPath + "/uniar.db"

		if _, err := os.Stat(uniarPath); err != nil {
			if err := os.Mkdir(uniarPath, 0755); err != nil {
				fmt.Println(err)
			}
		}

		if err := dbsetup(dbPath); err != nil {
			fmt.Println(err)
		}
		seed(dbPath)

		ctx := context.Background()
		db, err := repository.NewConnection()
		if err != nil {
			fmt.Println(err)
			return
		}
		q := repository.New()

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
				panic(err)
			}
			fmt.Println()
		}

		fmt.Printf("== 事務所ボーナスセットアップ ==\n")
		cob, _ := q.GetProducerOffice(ctx, db)
		ob := (&prompter.Prompter{
			Message: fmt.Sprintf("事務所ボーナス平均値 (現在値:%d) [1-17]", cob.Int64),
			Regexp:  regexp.MustCompile(`^([1-9]|1[0-7])$`),
		}).Prompt()
		obi, _ := strconv.Atoi(ob)
		if err := q.UpdateProducerOffice(ctx, db, sql.NullInt64{Int64: int64(obi), Valid: true}); err != nil {
			panic(err)
		}
		fmt.Println()

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
				h = "所持"
			}
			str := fmt.Sprintf("%s %s %s%s\n(現在: %s)", s.Photograph, s.Color, s.Member, ssrp, h)
			have := (&prompter.Prompter{
				Message: str,
				Choices: []string{"0", "1"},
				Default: "0",
			}).Prompt()
			hi, _ := strconv.Atoi(have)
			if err := q.UpdateProducerScene(ctx, db, repository.UpdateProducerSceneParams{
				Have: sql.NullInt64{Int64: int64(hi), Valid: true},
				ID:   s.ID,
			}); err != nil {
				panic(err)
			}
		}
	},
}

func seed(dbPath string) {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	result, err := db.ExecContext(ctx, string(mig_sql.Seed))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func dbsetup(dbPath string) error {
	config := database.Config{
		DbName: dbPath,
	}

	if options.DesiredFile == "" {
		f, err := os.CreateTemp("", "schema")
		if err != nil {
			return err
		}
		defer func() {
			filename := f.Name()
			if err := os.Remove(filename); err != nil {
				fmt.Println("一時ファイルの削除中にエラーが発生しました")
			}
		}()
		options.DesiredFile = f.Name()
		if _, err := f.Write(mig_sql.Schema); err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	options := &options
	db, err := sqlite3.NewDatabase(config)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("データベースの切断中にエラーが発生しました")
		}
	}()

	sqlParser := database.NewParser(parser.ParserModeSQLite3)
	sqldef.Run(schema.GeneratorModeSQLite3, db, sqlParser, options)

	return nil
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
