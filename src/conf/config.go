package conf

import (
	"github.com/spf13/viper"
)

var f, m *viper.Viper

func init() {
	f = viper.New()
	f.SetConfigType("yaml")
	f.SetConfigName("firebase")
	f.AddConfigPath("conf/environments/")

	m = viper.New()
	m.SetConfigType("yaml")
	m.SetConfigName("mysql")
	m.AddConfigPath("conf/environments/")
}

func GetFirebaseConfig() *viper.Viper {
	if err := f.ReadInConfig(); err != nil {
		return nil
	}
	return f
}

func GetMysqlConfig() *viper.Viper {
	if err := m.ReadInConfig(); err != nil {
		return nil
	}
	return m
}
