package config

import (
	"fmt"

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
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return &config, nil
}
