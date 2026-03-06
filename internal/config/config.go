// Package config
package config

import (
	"log"
	"os"
)

type Config struct {
	ServerConfig ServerConfig
	AIConfig     AIConfig
	DBConfig     DBConfig
}

type ServerConfig struct {
	Port string
}

type AIConfig struct {
	APIKey string
	Model  string
}

type DBConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBSSLMode  string
}

func NewConfig() *Config {
	aiCfg := AIConfig{
		APIKey: mustGetEnv("API_KEY"),
		Model:  getEnvOrDefault("MODEL", "gemini-3.1-flash-lite"),
	}

	serverCfg := ServerConfig{
		Port: getEnvOrDefault("PORT", "8080"),
	}

	dbCfg := DBConfig{
		DBName:     mustGetEnv("DB_NAME"),
		DBUser:     mustGetEnv("DB_USER"),
		DBPassword: mustGetEnv("DB_PASSWORD"),
		DBHost:     mustGetEnv("DB_HOST"),
		DBPort:     mustGetEnv("DB_PORT"),
		DBSSLMode:  mustGetEnv("DB_SSL_MODE"),
	}
	return &Config{
		ServerConfig: serverCfg,
		AIConfig:     aiCfg,
		DBConfig:     dbCfg,
	}
}

func mustGetEnv(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	log.Fatalf("unable to get the value of the %s", key)
	return ""
}

func getEnvOrDefault(key string, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
