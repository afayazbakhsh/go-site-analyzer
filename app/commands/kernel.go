package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "MyApp is a CLI application",
}

func Execute() {

	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfig() {

	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // directory of kernel (here)
	viper.SetDefault("crawler.maxWorker", 2)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("⚠️ No config file found, using defaults")
	} else {
		fmt.Println("✅ Using config file:", viper.ConfigFileUsed())
	}
}
