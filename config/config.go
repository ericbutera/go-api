// Provides configuration for the API
package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Application Configuration
type Config struct {
	// Enable DataDog integration
	// DataDog                 bool   `mapstructure:"data_dog"`
	// DataDogApiKey           string `mapstructure:"data_dog_api_key"`
	Port                    string `mapstructure:"port"`
	AppName                 string `mapstructure:"app_name"`
	ServiceName             string `mapstructure:"service_name"`
	Env                     string `mapstructure:"env"`
	Version                 string `mapstructure:"version"`
	JaegerCollectorEndpoint string `mapstructure:"jaeger_collector_endpoint"`
}

func NewAppConfig(path *string, file *string) (Config, error) {
	viper.AddConfigPath(*path)
	viper.SetConfigName(*file)
	viper.SetConfigType("yaml")

	viper.SetDefault("jaeger", false)
	viper.SetDefault("jaeger_collector_endpoint", "http://localhost:14268/api/traces")
	// viper.SetDefault("data_dog", false)
	// viper.SetDefault("data_dog_api_key", "")
	viper.SetDefault("app_name", "go-api")
	viper.SetDefault("service_name", "go-api")
	viper.SetDefault("env", "dev")
	viper.SetDefault("version", "0.0.1")

	read_err := viper.ReadInConfig()
	if read_err != nil {
		log.Print(fmt.Errorf("fatal error config file: %w", read_err))
	}

	var config Config
	parse_err := viper.Unmarshal(&config)
	if parse_err != nil {
		log.Print(fmt.Errorf("cannot parse config %s", parse_err))
	}

	log.Printf("config %+v", config)
	return config, parse_err
}
