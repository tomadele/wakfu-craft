package database

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/wakfu-craft/server/database"
)

var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes database with Wakfu JSON files",
	Long:  "Database related commands such as update.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := database.Sync(); err != nil {
			log.Fatalf("Failed to sync database: %v", err)
		}
	},
}
