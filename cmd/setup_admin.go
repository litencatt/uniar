package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

var setupAdminCmd = &cobra.Command{
	Use:   "admin [producer_id]",
	Short: "Set producer as admin",
	Long:  `Set specified producer as admin user.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setupAdmin(args[0])
	},
}

func init() {
	setupCmd.AddCommand(setupAdminCmd)
}

func setupAdmin(producerIDStr string) {
	producerID, err := strconv.ParseInt(producerIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid producer ID: %v", err)
	}

	dbPath := GetDbPath()
	database, err := repository.NewConnection(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	ctx := context.Background()

	// Update producer to admin
	_, err = database.ExecContext(ctx, "UPDATE producers SET is_admin = 1 WHERE id = ?", producerID)
	if err != nil {
		log.Fatalf("Failed to set admin: %v", err)
	}

	fmt.Printf("Producer ID %d has been set as admin\n", producerID)
}