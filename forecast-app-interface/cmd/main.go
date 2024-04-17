package main

import (
	"forecast-app-interface/config"
	"forecast-app-interface/internal/app"
	"log"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("Cannot parse configs: %s\n", err)
	}

	app.Run(cfg)
}
