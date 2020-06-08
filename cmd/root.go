package cmd

import (
	"log"

	"github.com/pokeman-service/api"

	"github.com/pokeman-service/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "example",
		Run: run,
	}

	rootCmd.Flags().IntP("app.port", "p", 8080, "Port which the http server runs on")
	rootCmd.Flags().StringP("app.profile", "e", "e1", "Config Env Profile")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	initConfig(cmd)
	api.InitWebServer()
}

func initConfig(cmd *cobra.Command) {

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
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
