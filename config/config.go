package config

import (
	"errors"
	"os"

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

type Config struct {
	ApiConfig
	DbConfig
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
