package config

import (
	"github.com/spf13/viper"
)

type DbConfig struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Port     string `mapstructure:"port"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("env")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
