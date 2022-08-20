package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/mysql"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/litencatt/unisonair/repository"
	mig_sql "github.com/litencatt/unisonair/sql"
	"github.com/xo/dburl"
)

var (
	options  = sqldef.Options{}
	skipSeed bool
)

func run(dsn string) error {
	ctx := context.Background()
	db, err := dburl.Open(dsn)
	if err != nil {
		return err
	}

	queries := repository.New(db)

	// list all authors
	groups, err := queries.GetGroup(ctx)
	if err != nil {
		return err
	}
	log.Println(groups)

	return nil
}

func migrate(dsn string) error {
	u, err := dburl.Parse(dsn)
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
				fmt.Errorf("一時ファイルの削除中にエラーが発生しました")
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
	db, err := mysql.NewDatabase(config)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Errorf("データベースの切断中にエラーが発生しました")
		}
	}()

	sqlParser := database.NewParser(parser.ParserModeMysql)
	sqldef.Run(schema.GeneratorModeMysql, db, sqlParser, options)

	return nil
}

func main() {
	dsn := "mysql://root@db:3306/unisonair?parseTime=true&loc=Asia%2FTokyo"

	if err := migrate(dsn); err != nil {
		log.Fatal(err)
	}
	if err := run(dsn); err != nil {
		log.Fatal(err)
	}
}
