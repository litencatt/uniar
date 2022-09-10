/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"os"

	"github.com/litencatt/unisonair/repository"
	"github.com/spf13/cobra"
	"github.com/xo/dburl"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		dsn := os.Getenv("UNIAR_DSN")
		db, err := dburl.Open(dsn)
		if err != nil {
			log.Print(err)
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
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
