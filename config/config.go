package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"

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

type TokenConfig struct {
	IssuerName       string                 `json:"issuerName"`
	JwtSecretKey     []byte                 `json:"jwtSecretKey"`
	JwtSigningMethod *jwt.SigningMethodHMAC `json:"jwtSigningMethod"`
	JwtExpiredTime   time.Duration          `json:"jwtExpiredTime"`
}

type Config struct {
	DbConfig
	ApiConfig
	TokenConfig
}

func (c *Config) ConfigurationDB() error {
	err := godotenv.Load()
	if err != nil {
		return err
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
	tokenExpired, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME"))
	if err != nil {
		return err
	}

	// config token
	c.TokenConfig = TokenConfig{
		IssuerName:       os.Getenv("ISSUER_NAME"),
		JwtSecretKey:     []byte(os.Getenv("JWT_SECRET_KEY")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtExpiredTime:   time.Duration(tokenExpired) * time.Hour,
	}

	if c.Host == "" || c.Port == "" || c.User == "" || c.Password == "" || c.Database == "" || c.Driver == "" || c.ApiPort == "" || c.IssuerName == "" || c.JwtSecretKey == nil || c.JwtSigningMethod == nil || tokenExpired == 0 || c.JwtExpiredTime == 0 {
		return fmt.Errorf("missing required env variables")
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
