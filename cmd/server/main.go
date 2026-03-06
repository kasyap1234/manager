package main

import (
	"log"
	"manager/internal/config"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

}
