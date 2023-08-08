package conf

import (
	"github.com/spf13/viper"
)

var m *viper.Viper

func GetMysqlConfig() *viper.Viper {

	m = viper.New()
	m.SetConfigType("yaml")
	m.SetConfigName("mysql")
	m.AddConfigPath("config/")

	if err := m.ReadInConfig(); err != nil {
		return nil
	}
	return m
}
