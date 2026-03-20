// Package config
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	ConnStr string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("No .env file found: %v", err)
	}
	aiCfg := AIConfig{
		APIKey: mustGetEnv("GEMINI_API_KEY"),
		Model:  getEnvOrDefault("MODEL", "gemini-3.1-flash-lite-preview"),
	}

	serverCfg := ServerConfig{
		Port: getEnvOrDefault("PORT", "8080"),
	}

	dbCfg := DBConfig{
		ConnStr: mustGetEnv("DB_CONN_STR"),
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
