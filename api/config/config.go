package config

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Define the config yml structure
type Config struct {
	Cookie ConfigCookie `yaml:"cookie"`
	HTTP   ConfigHTTP   `yaml:"http"`
	DB     ConfigDB     `yaml:"db"`
}

type ConfigCookie struct {
	Secure bool `yaml:"secure"`
}

type ConfigHTTP struct {
	Port int `yaml:"port"`
}

type ConfigDB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

// Initialise the config from file
func New() (*Config, error) {
	var filePath string
	config := &Config{}

	flag.StringVar(&filePath, "config", "./config.yml", "path to config file")
	flag.Parse()

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return config, nil
}
