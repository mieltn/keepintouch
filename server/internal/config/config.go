package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecretKey           string
	JWTAccessLifetimeMins  int
	JWTRefreshLifetimeMins int
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	JWTAccessLifetimeMins, err := getInt("JWT_ACCESS_LIFETIME_MINS")
	if err != nil {
		return nil, err
	}
	JWTRefreshLifetimeMins, err := getInt("JWT_REFRESH_LIFETIME_MINS")
	if err != nil {
		return nil, err
	}

	return &Config{
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		JWTAccessLifetimeMins: JWTAccessLifetimeMins,
		JWTRefreshLifetimeMins: JWTRefreshLifetimeMins,
	}, nil
}

func getInt(key string) (int, error) {
	strVal := os.Getenv(key)
	val, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, err
	}
	return val, nil
}
