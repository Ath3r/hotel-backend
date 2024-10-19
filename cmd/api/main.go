package main

import (
	"log"

	"github.com/Ath3r/hotel-backend/cmd/api/server"
	"github.com/Ath3r/hotel-backend/internal/config"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app, err := server.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}
}
