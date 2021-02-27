package config

import (
	"os"

	"github.com/spf13/viper"
)

// EnableTestMode sets test configuration as current
func EnableTestMode() {
	viper.SetDefault("db", ":memory:")
	viper.SetDefault("host", ":8080")
}

// EnableDefaultMode reads config.json file or uses default configuration
func EnableDefaultMode() error {
	if os.Getenv("ENVIRONMENT") == "production" {
		return enableProductionMode()
	}
	return enableDevelopmentMode()
}

// enableDevelopmentMode sets dev configuration as current
func enableDevelopmentMode() error {
	viper.SetConfigType("json")
	viper.SetDefault("db", ".dev/kronos.db")
	viper.SetDefault("host", ":8080")
	viper.AddConfigPath(".dev/")
	err := viper.ReadInConfig()
	return err
}

// enableProductionMode sets production configuration as current
func enableProductionMode() error {
	viper.SetConfigType("json")
	viper.SetDefault("db", "/usr/local/var/kronos/kronos.db")
	viper.SetDefault("host", ":8080")
	viper.AddConfigPath("/usr/local/etc/kronos/")
	err := viper.ReadInConfig()
	return err
}

// GetString gets a configuration value in string format
func GetString(name string) string {
	return viper.GetString(name)
}
