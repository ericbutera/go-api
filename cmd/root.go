package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-api",
		Short: "go-api",
		Long:  "go-api k8s and golang playground",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("env", "e", "dev", "environment")
	viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))

	rootCmd.AddCommand(NewServerCmd())
}

func initConfig() {
	// this allows NISTY_ENV to work automagically!
	// https://github.com/spf13/viper#env-example
	viper.SetEnvPrefix("goapi")
	viper.AutomaticEnv()
}
