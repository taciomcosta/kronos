package config

var defaultConfig map[string]string = map[string]string{
	"db":   "kronos.db",
	"host": ":8080",
}

var testConfig map[string]string = map[string]string{
	"db":   ":memory:",
	"host": ":8080",
}

var currentConfig map[string]string = defaultConfig

// EnableTestMode sets test configuration as current
func EnableTestMode() {
	currentConfig = testConfig
}

// GetString gets a configuration value in string format
func GetString(name string) string {
	value := currentConfig[name]
	return value
}
