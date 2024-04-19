package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"

	"forecast-app-auth/config"
	"forecast-app-auth/internal/app"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("Cannot parse configs: %s\n", err)
	}

	app.Run(cfg)
}
