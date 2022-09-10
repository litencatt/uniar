package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/sqldef"
	"github.com/litencatt/unisonair/repository"
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

func seed(dsn string) error {
	ctx := context.Background()
	db, err := dburl.Open(dsn)
	if err != nil {
		return err
	}
	queries := repository.New(db)
	queries.SeedGroups(ctx)
	queries.SeedCenterSkills(ctx)
	queries.SeedColorTypes(ctx)
	queries.SeedLives(ctx)
	queries.SeedMembers(ctx)
	queries.SeedMusic(ctx)
	queries.SeedPhotograph(ctx)
	queries.SeedScenes(ctx)
	queries.SeedSkills(ctx)
	return nil
}

func main() {
	dsn := "mysql://root@db:3306/unisonair?parseTime=true&loc=Asia%2FTokyo"
	if err := seed(dsn); err != nil {
		log.Fatal(err)
	}
	if err := run(dsn); err != nil {
		log.Fatal(err)
	}
}
