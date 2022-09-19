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

	_ "github.com/mattn/go-sqlite3"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/sqlite3"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
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
			fmt.Print(err)
		}
		seed(dbPath)
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
