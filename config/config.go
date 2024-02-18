package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Driver   string
}

type LogConfig struct {
	FilePath string
}

type JWTConfig struct {
	SecretKey string
	Lifetime  time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	LogConfig
	JWTConfig
}

/*
Validate env configuration value based on configuration Struct
*/
func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.LogConfig = LogConfig{
		FilePath: os.Getenv("LOG_FILE_PATH"),
	}

	tokenLifeTime, err := time.ParseDuration(os.Getenv("JWT_LIFETIME"))
	if err != nil {
		return err
	}

	c.JWTConfig = JWTConfig{
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
		Lifetime:  tokenLifeTime,
	}

	// TODO: Validate config
	if c.ApiConfig.ApiPort == "" {
		return errors.New("ApiPort cannot be empty")
	}

	return nil
}

/*
Call read config
*/
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
