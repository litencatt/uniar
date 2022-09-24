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

	"github.com/litencatt/uniar/repository"
	mig_sql "github.com/litencatt/uniar/sql"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/sqlite3"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/spf13/cobra"
)

var (
	options = sqldef.Options{}
)

var setupMigrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		dbPath := GetDbPath()
		db, err := repository.NewConnection(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		setupMkdir()
		setupMigrate(dbPath)
		setupSeed(ctx, db, dbPath)
	},
}

func setupMkdir() {
	uniarPath := GetUniarPath()
	if _, err := os.Stat(uniarPath); err != nil {
		if err := os.Mkdir(uniarPath, 0750); err != nil {
			fmt.Println(err)
		}
	}
}

func setupMigrate(dbPath string) error {
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

func setupSeed(ctx context.Context, db *sql.DB, dbPath string) {
	result, err := db.ExecContext(ctx, string(mig_sql.Seed))
	if err != nil {
		fmt.Println(err)
	}
	if _, ok := os.LookupEnv("DEBUG"); ok {
		fmt.Println(result)
	}
}

func init() {
	setupCmd.AddCommand(setupMigrateCmd)
}
