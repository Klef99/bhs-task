package main

import (
	"log"

	"github.com/Klef99/bhs-task/config"
	"github.com/Klef99/bhs-task/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %w", err)
	}

	// Run
	app.Run(cfg)
}
