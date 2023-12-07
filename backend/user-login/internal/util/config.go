package util

import (
	"github.com/spf13/viper"
	"os"
)

var c *config

func init() {
	var err error
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err = viper.Unmarshal(&c); err != nil {
		panic(err)
	}
}

func GetConfig() *config {
	return c
}

type config struct {
	SERVER struct {
		SERVER_PORT string `mapstructure:"port"`
	} `mapstructure:"server"`
}
