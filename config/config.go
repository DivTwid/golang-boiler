package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DatabasePostgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database_postgres"`

	DatabaseMysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database_mysql"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	App struct {
		Port     string `yaml:"port"`
		LogLevel string `yaml:"log_level"`
	} `yaml:"app"`
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	configFile := "config." + env + ".yaml"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config map[string]Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	envConfig, ok := config[env]
	if !ok {
		return nil, fmt.Errorf("configuration for environment '%s' not found", env)
	}

	return &envConfig, nil
}
