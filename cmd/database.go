package cmd

import (
	"github.com/spf13/cobra"

	"github.com/wakfu-craft/cmd/database"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database related commands.",
	Long:  "Database related commands such as init or sync.",
}

func init() {
	databaseCmd.AddCommand(database.InitCmd)
	databaseCmd.AddCommand(database.SyncCmd)
	rootCmd.AddCommand(databaseCmd)
}
