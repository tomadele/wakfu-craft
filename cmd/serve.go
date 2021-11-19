package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/wakfu-craft/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the HTTP server",
	Long:  "Launches the HTTP server and exposes its API",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.New(); err != nil {
			log.Fatalf("Failed to launch server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
