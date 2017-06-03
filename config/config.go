package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ConfigReader interface defines methods to get the configurations from yml files.
type ConfigReader interface {
	AutomaticEnv()
	GetString(key string) string
	GetInt(key string) int
	SetConfigFile(in string)
	ReadInConfig() error
}

// MainConfig implements the Config interface to get configuration from yml files.
type MainConfig struct {
	viper *viper.Viper
}

// New create and instance of MainConfig struct
func New() *MainConfig {
	return &MainConfig{
		viper: viper.New(),
	}
}

// AutomaticEnv check ENV variables for all. keys set in config, default & flags
func (config *MainConfig) AutomaticEnv() {
	config.viper.AutomaticEnv()
}

// GetString look for the given key to return a string
func (config *MainConfig) GetString(key string) string {
	return config.viper.GetString(key)
}

// GetInt look for the given key to return a int
func (config *MainConfig) GetInt(key string) int {
	return config.viper.GetInt(key)
}

// SetConfigFile set the file from which read configs
func (config *MainConfig) SetConfigFile(in string) {
	config.viper.SetConfigFile(in)
}

// ReadInConfig will discover and load the configuration file from disk
func (config *MainConfig) ReadInConfig() error {
	return config.viper.ReadInConfig()
}

// Init read config from file
func Init() (*MainConfig, error) {
	config := New()

	goEnv := os.Getenv("GOENV")
	config.SetConfigFile(fmt.Sprintf("%s/.uniplaces/%s.yaml", os.Getenv("HOME"), goEnv))

	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error reading config file: %s", err)
	}

	return config, nil
}
