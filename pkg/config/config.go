package config

import "os"

type Config struct {
	Database     *DatabaseConfig
	CookieSecret string
}

type DatabaseConfig struct {
	Host        string
	Port        string
	User        string
	Password    string
	Database    string
	LoggerLevel string
}

func DBConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		User:        os.Getenv("DB_USER"),
		Password:    os.Getenv("DB_PASSWORD"),
		Database:    os.Getenv("DB_DATABASE"),
		LoggerLevel: os.Getenv("DB_LOGGER_LEVEL"),
	}
}

func New() *Config {
	return &Config{
		Database: DBConfig(),
	}
}
