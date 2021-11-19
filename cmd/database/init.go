package database

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/wakfu-craft/server/database"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes database",
	Long:  "Initializes database",
	Run: func(cmd *cobra.Command, args []string) {
		conf := database.Configuration{
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.db"),
		}

		database.Initialize(conf)
	},
}
