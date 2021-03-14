package config

import (
	"os"

	"github.com/spf13/viper"
)

var defaultConfigPath string
var defaultDBPath string

// EnableTestMode sets test configuration as current
func EnableTestMode() {
	viper.SetDefault("db", ":memory:")
	viper.SetDefault("host", ":8080")
}

// EnableDefaultMode reads config.json file or uses default configuration
func EnableDefaultMode() error {
	if os.Getenv("ENVIRONMENT") == "development" {
		return enableDevelopmentMode()
	}
	return enableProductionMode()
}

func enableDevelopmentMode() error {
	viper.SetConfigType("json")
	viper.SetDefault("db", ".dev/kronos.db")
	viper.SetDefault("host", ":8080")
	viper.AddConfigPath(".dev/")
	err := viper.ReadInConfig()
	return err
}

func enableProductionMode() error {
	viper.SetConfigType("json")
	viper.SetDefault("db", defaultDBPath)
	viper.SetDefault("host", ":8080")
	viper.AddConfigPath(defaultConfigPath)
	err := viper.ReadInConfig()
	return err
}

// GetString gets a configuration value in string format
func GetString(name string) string {
	return viper.GetString(name)
}
