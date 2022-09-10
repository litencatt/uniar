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
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/mysql"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/litencatt/unisonair/sql"
	"github.com/spf13/cobra"
	"github.com/xo/dburl"
)

var (
	options = sqldef.Options{}
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := execute(); err != nil {
			log.Print(err)
		}
	},
}

func execute() error {
	dsn := os.Getenv("UNIAR_DSN")
	u, _ := dburl.Parse(dsn)
	password, _ := u.User.Password()
	port, _ := strconv.Atoi(u.Port())
	config := database.Config{
		DbName:   strings.TrimPrefix(u.Path, "/"),
		User:     u.User.Username(),
		Password: password,
		Host:     u.Hostname(),
		Port:     port,
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
		if _, err := f.Write(sql.Schema); err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	options := &options
	db, err := mysql.NewDatabase(config)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("データベースの切断中にエラーが発生しました")
		}
	}()

	sqlParser := database.NewParser(parser.ParserModeMysql)
	sqldef.Run(schema.GeneratorModeMysql, db, sqlParser, options)

	return nil
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
