package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Port     string
	Origin   string
	Database string
}

var cfg *config

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("config initialization error: %v", err)
	}

	cfg = &config{
		Port:     viper.GetString("server.port"),
		Origin:   viper.GetString("server.origin"),
		Database: viper.GetString("server.database"),
	}

	return nil
}

func GetConfig() *config {
	return cfg
}
