package config

import (
	"github.com/spf13/viper"
)

const configPath = "config/config.yaml"

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Database DatabaseConfig `yaml:"database"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig
	MongoDB  MongoDBConfig
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type MongoDBConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
	Bucket   string `yaml:"bucket"`
}

func NewConfig() (*Config, error) {
	var err error
	var config Config

	viper.SetConfigFile(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
