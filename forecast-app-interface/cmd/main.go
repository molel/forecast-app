package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"forecast-app-interface/config"
	"forecast-app-interface/internal/app"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("Cannot parse configs: %s\n", err)
	}

	app.Run(cfg)
}
