package config

import "os"

type Config struct {
	Database *DatabaseConfig
	Cors     *CorsConfig
}

type DatabaseConfig struct {
	Host        string
	Port        string
	User        string
	Password    string
	Database    string
	LoggerLevel string
}

type CorsConfig struct {
	AllowedOrigins string
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

func CORSConfig() *CorsConfig {
	return &CorsConfig{
		AllowedOrigins: os.Getenv("CORS_ALLOWED_ORIGINS"),
	}
}

func New() *Config {
	return &Config{
		Database: DBConfig(),
		Cors:     CORSConfig(),
	}
}
