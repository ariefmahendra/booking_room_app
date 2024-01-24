package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}

type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) ConfigurationDB() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file")
	}

	// config db
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
	// config PORT app
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("PORT"),
	}
	// another config

	if c.Host == "" || c.Port == "" || c.User == "" || c.Password == "" || c.Database == "" || c.Driver == "" || c.ApiPort == "" {
		return fmt.Errorf("missing required environmet varible")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ConfigurationDB(); err != nil {
		return nil, err
	}
	return cfg, nil
}
