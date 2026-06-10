package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port string
	Host string
}

var ENV = NewConfig()

func NewConfig() Config {
	return Config{
		Host: getEnv("HOST", "0.0.0.0"),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func getIntEnv(key string, fallback int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
