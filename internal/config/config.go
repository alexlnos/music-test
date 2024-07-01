package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	CheckInterval time.Duration
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	println(time.Duration(viper.GetInt("checkInterval")))

	return &Config{
		CheckInterval: time.Duration(viper.GetInt("checkInterval")) * time.Second,
	}
}
