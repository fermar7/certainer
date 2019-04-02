package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func initConfig(path string) (certainerConfig, error) {
	viper.SetConfigType("toml")
	viper.SetConfigFile("certainer.toml")

	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return certainerConfig{}, fmt.Errorf("Error reading config: %s", err)
	}

	var config certainerConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		return certainerConfig{}, fmt.Errorf("Error parsing config: %s", err)
	}

	return config, nil
}

type certainerConfig struct {
	Authority string         `mapstructure:"authority"`
	Database  databaseConfig `mapstructure:"database"`
	Ports     portConfig     `mapstructure:"ports"`
}

type databaseConfig struct {
	Path     string `mapstructure:"path"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type portConfig struct {
	WellKnown int64 `mapstructure:"wellknown"`
	API       int64 `mapstructure:"api"`
}
