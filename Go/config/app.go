package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppName        string          `mapstructure:"app_name"`
	ServerConfig   *ServerConfig   `mapstructure:"server"`
	DatabaseConfig *DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Address string `mapstructure:"address"`
	Port    uint16 `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}

func LoadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
