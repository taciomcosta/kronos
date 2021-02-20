package config

import (
	"github.com/spf13/viper"
)

// EnableDefaultMode reads config.json file or uses default configuration
func EnableDefaultMode() error {
	viper.SetConfigType("json")
	viper.SetDefault("db", "kronos.db")
	viper.SetDefault("host", ":8080")
	viper.AddConfigPath(getConfigFilePath())
	err := viper.ReadInConfig()
	return err
}

func getConfigFilePath() string {
	return "/etc/kronos/"
}

// EnableTestMode sets test configuration as current
func EnableTestMode() {
	viper.SetDefault("db", ":memory:")
	viper.SetDefault("host", ":8080")
}

// GetString gets a configuration value in string format
func GetString(name string) string {
	return viper.GetString(name)
}
