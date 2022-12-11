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

	"github.com/Songmu/prompter"
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
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := setupMkdir(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if prompter.YN("do you execute migration and seed?", false) {
			if err := migrate(ctx, db, dbPath); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			fmt.Println("skip migration")
		}
	},
}

func migrate(ctx context.Context, db *sql.DB, dbPath string) error {
	if err := setupMigrate(dbPath); err != nil {
		return err
	}
	if err := setupSeed(ctx, db, dbPath); err != nil {
		return err
	}

	return nil
}

func setupMkdir() error {
	uniarPath := GetUniarPath()
	if _, err := os.Stat(uniarPath); err != nil {
		if err := os.Mkdir(uniarPath, 0750); err != nil {
			return err
		}
	}
	return nil
}

func setupMigrate(dbPath string) error {
	fmt.Println("start migration")
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
	// Run with hide schema migration diff temporary
	tmp := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	sqldef.Run(schema.GeneratorModeSQLite3, db, sqlParser, options)
	os.Stdout = tmp

	fmt.Println("end migration")
	return nil
}

func setupSeed(ctx context.Context, db *sql.DB, dbPath string) error {
	fmt.Println("start seed")
	result, err := db.ExecContext(ctx, string(mig_sql.Seed))
	if err != nil {
		return err
	}
	if _, ok := os.LookupEnv("DEBUG"); ok {
		fmt.Println(result)
	}
	fmt.Println("end seed")
	return nil
}

func init() {
	setupCmd.AddCommand(setupMigrateCmd)
}
