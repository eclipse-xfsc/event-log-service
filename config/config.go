package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type logServiceConfig struct {
	Keywords []string `mapstructure:"keywords"`
}

var CurrentLogServiceConfig logServiceConfig

func LoadConfig() {
	readConfig()
	if err := viper.Unmarshal(&CurrentLogServiceConfig); err != nil {
		panic(err)
	}
}

func readConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("LOGSERVICE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Printf("Configuration not found but environment variables will be taken into account.")
			viper.AutomaticEnv()
		} else {
			panic(err)
		}
	}
}
