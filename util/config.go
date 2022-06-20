package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	SECRET_NAME  string `mapstructure:"SECRET_NAME"`
	SECRET_VALUE string `mapstructure:"SECRET_VALUE"`
}

func LoadConfig(path string) (config Config, err error) {
	// add env file configuration
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// add environment variables - priority
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Config file found and successfully parsed
	err = viper.Unmarshal(&config)
	return config, err
}
