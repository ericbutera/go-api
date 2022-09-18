// Provides configuration for the API
package config

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
	UseOpenTel              bool   `mapstructure:"use_opentel"`
}
