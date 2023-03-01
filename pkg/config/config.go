package config

import (
	"github.com/spf13/viper"
	"log"
)

func GetConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configurations")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config error : ", err.Error())
	}
}
