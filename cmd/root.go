package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/pokeman-service/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use: "example",
	}
)

//Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initCobraConfigFlags() {

}

func initConfig() {
	rootCmd.Flags().IntP("app.port", "p", 8080, "Port which the http server runs on")
	rootCmd.Flags().StringP("app.profile", "e", "e1", "Config Env Profile")

	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Fatalf("Failed to Bind Command Line Flags, %s", err)
	}

	viper.SetConfigName(viper.GetString("app.profile"))
	viper.SetConfigType("yml")
	viper.AddConfigPath("resources")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error Reading Config File, %s", err)
	}

	if err := viper.Unmarshal(&model.Config); err != nil {
		log.Fatalf("Error Serializing Config, %s", err)
	}

	log.Print("Config ", model.Config)

}
