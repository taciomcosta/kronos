// +build linux

package config

func init() {
	defaultConfigPath = "/etc/kronos/"
	defaultDBPath = "/var/lib/kronos/kronos.db"
}
