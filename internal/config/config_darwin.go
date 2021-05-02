// +build darwin

package config

func init() {
	defaultConfigPath = "/usr/local/etc/kronos/"
	defaultDBPath = "/usr/local/var/kronos.db"
	defaultHost = ":8080"
}
