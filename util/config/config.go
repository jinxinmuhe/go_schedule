package config

import "github.com/spf13/viper"

var Viper *viper.Viper

func InitConfig(path string) {
	Viper = viper.New()
	if path == "" {
		Viper.SetConfigFile("../../conf/app.toml")
	} else {
		Viper.SetConfigFile(path)
	}
	if err := Viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
