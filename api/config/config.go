package config

import (
	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

// Define the config yml structure
type Config struct {
	IexApiKey  string `mapstructure:"IEXCLOUD_API_KEY"`
	JwtKey     string `mapstructure:"JWT_KEY"`
	DbHost     string `mapstructure:"DB_HOST_NAME"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUserName string `mapstructure:"DB_USER_NAME"`
	DbUserPw   string `mapstructure:"DB_USER_PW"`
	DbPort     uint   `mapstructure:"DB_PORT"`
	AppPort    uint   `mapstructure:"APP_PORT"`
}

// Initialise the config from file
func New() (*Config, error) {
	var config Config

	config.AppPort = 8080
	config.DbHost = "db"
	config.DbName = "dev"
	config.DbPort = 5432
	config.DbUserName = "dev"
	config.DbUserPw = "dev"
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("JWT_KEY", "VerySecretKey")

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return &config, err
	} else {
		err := viper.Unmarshal(&config)
		if err != nil {
			errors.Wrap(err, "could not unmarshal config")
		}
	}

	return &config, nil
}
