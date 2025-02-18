package config

import (
	"github.com/spf13/viper"
)

// use -ldflags flag to change value of these variables
//
// example: -ldflags="-X github.com/binarstrike/chiya-chatbot-go/config.APP_VERSION=1.0.4-beta"
var (
	APP_ENV     = "development"
	APP_VERSION = "0.0.1-alpha"
)

type databaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string `mapstructure:"db_name"`
}

type redisConfig struct {
	Host string
	Port int
}

type botConfig struct {
	AppId string `mapstructure:"app_id"`
	Token string
}

type Config struct {
	Database databaseConfig `mapstructure:"database"`
	Redis    redisConfig    `mapstructure:"redis"`
	Bot      botConfig      `mapstructure:"bot"`
}

func InitConfig() (*Config, error) {
	config := new(Config)

	v := viper.New()

	v.AddConfigPath(".") // lokasi file konfigurasi
	v.AddConfigPath("config")
	v.SetConfigType("toml")   // tipe/format file konfigurasi
	v.SetConfigName("config") // nama file konfigurasi tanpa nama ekstensi
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
