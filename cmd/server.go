package main

import (
	"log"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/yazmeyaa/chat_backend/internal/services/config"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config := config.New()

	slog.Info(config.JWT.Key)
}
