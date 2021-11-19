package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile  string
	logLevel string

	rootCmd = &cobra.Command{
		Use:   "wakfu-craft",
		Short: "WakfuCraft is a web application allowing users to help themselves craft in Wakfu.",
		Long:  "WakfuCraft is a web application allowing users to help themselves craft in Wakfu.",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(os.Stderr, err)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is config/serverTemp.yml)")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "log level (default is info)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.Getwd()
		cobra.CheckErr(err)

		// Search default "serverTemp.yml" config file
		viper.AddConfigPath(pwd + "/config")
		viper.SetConfigType("yml")
		viper.SetConfigName("server")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initLogger() {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Failed to parse log level: %v", err)
	}
	log.SetLevel(level)
}
