// Package config
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AIConfig AIConfig
	DBConfig DBConfig
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

func NewConfig()*Config{
	aiCfg :=AIConfig{
		APIKey: mustGetEnv("API_KEY"),
		Model: getEnvOrDefault("MODEL", "gemini-3.1-flash-lite"),
	}
	dbCfg :=DBConfig{
		DBName: mustGetEnv("DB_NAME"),
		
	}
	return &Config{
		aiCfg,
		dbCfg,
	}
}


func mustGetEnv(key string)string{
	if val :=os.Getenv(key); val !=""{
		return val
	}
	return fmt.Errorf("unable to get the value of the %s",key)
}

func getEnvOrDefault(key string,defaultValue string)string{
	if val :=os.Getenv(key); val !=""{
		return val
	}
	return defaultValue
}
