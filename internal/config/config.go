package config

import (
	"github.com/Ath3r/hotel-backend/internal/constants"
	"github.com/spf13/viper"
)

type Config struct {
	Port             int    `mapstructure:"PORT"`
	Environment      string `mapstructure:"ENVIRONMENT"`
	Debug            bool   `mapstructure:"DEBUG"`

	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
}

var AppConfig *Config

func LoadConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("internal/config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return err
	}

	if AppConfig.Port == 0 || AppConfig.Environment == "" || AppConfig.DatabaseHost == "" || AppConfig.DatabasePort == "" || AppConfig.DatabaseName == "" {
		return constants.ErrEmptyVar
	}

	return nil
}

