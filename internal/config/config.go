package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	APIKey  string
	BaseURL string
}

// LoadConfig loads the configuration from file and environment variables
func LoadConfig() (*Config, error) {
	viper.SetConfigName(".cyverApiCli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	config := &Config{
		APIKey:  viper.GetString("api_key"),
		BaseURL: viper.GetString("base_url"),
	}

	return config, nil
} 