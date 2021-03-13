package config

import (
	"github.com/spf13/viper"
)

var defaultConfigPath = "./dev/kronos.db"
var defaultDBPath = ".dev/"

// EnableTestMode sets test configuration as current
func EnableTestMode() {
	viper.SetDefault("db", ":memory:")
	viper.SetDefault("host", ":8080")
}

// EnableDefaultMode reads config.json file or uses default configuration
func EnableDefaultMode() error {
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
