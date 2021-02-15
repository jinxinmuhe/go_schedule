package config

import (
	"go_schedule/util/log"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

func InitConfig(path string) error {
	Viper = viper.New()
	if path == "" {
		Viper.SetConfigFile("../../conf/app.toml")
	} else {
		Viper.SetConfigFile(path)
	}
	if err := Viper.ReadInConfig(); err != nil {
		log.ErrLogger.Printf("init config error:%+v", err)
		return err
	}
	return nil
}
