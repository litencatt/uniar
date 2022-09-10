/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
