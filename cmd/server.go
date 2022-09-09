// TODO: look into sub commands
package cmd

import (
	"log"
	"github.com/ericbutera/go-api/api"
	"github.com/ericbutera/go-api/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewConfig() *config.Config {
	viper.SetDefault("Env", "dev")

	var config config.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	log.Printf("config: %+v", config)

	return &config
}

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Long:  "Run api server",
		Run:   RunServer,
	}
	return cmd
}

func RunServer(cmd *cobra.Command, args []string) {
	config := NewConfig()
	app, err := api.NewApp(config)
	if err != nil {
		panic(err)
	}

	app.Serve()
}
