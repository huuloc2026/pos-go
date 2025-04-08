package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // type of config file
	viper.AddConfigPath("./config") // path to look for config file

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
}
