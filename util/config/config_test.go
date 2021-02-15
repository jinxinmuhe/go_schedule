package config

import "testing"

func TestViperConfig(t *testing.T) {
	conn := Viper.Get("mongodb.conn")
	list := Viper.GetStringSlice("zookeeper.hosts")
	t.Log(conn)
	t.Log(list)
}