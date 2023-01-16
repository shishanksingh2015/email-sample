package config

import (
	"github.com/labstack/gommon/log"
	"github.com/shishanksingh2015/email-sample/model"
	"github.com/spf13/viper"
	"strings"
)

func ReadConfig(path string) model.ServerConfig {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load config err: ", err)
	}
	config := model.ServerConfig{}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to parse config err: ", err)
	}
	return config
}
