package config

import (
	"flag"
	"os"
)

type Config struct {
	Address string
	BaseURL string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func GetConfig() *Config {

	serverAddress := getEnv("SERVER_ADDRESS", "localhost:8080")
	baseURL := getEnv("BASE_URL", "http://localhost:8080/")

	flag.StringVar(&serverAddress, "a", serverAddress, "Server address")
	flag.StringVar(&baseURL, "b", baseURL, "Base URL")

	flag.Parse()

	return &Config{
		Address: serverAddress,
		BaseURL: baseURL,
	}
}
