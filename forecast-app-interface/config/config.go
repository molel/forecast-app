package config

import (
	"flag"
	"fmt"
)

type Config struct {
	HTTP                  uint64
	AuthServiceAddress    string
	PredictServiceAddress string
}

func Parse() (*Config, error) {
	cfg := &Config{}

	flag.Uint64Var(&cfg.HTTP, "http", 8080, "listened by app port")
	flag.StringVar(&cfg.AuthServiceAddress, "auth-service-address", "", "http address auth service is being hosted on")
	flag.StringVar(&cfg.PredictServiceAddress, "predict-service-address", "", "http address predict service is being hosted on")
	flag.Parse()

	if flag.NArg() > 0 {
		return nil, fmt.Errorf("invalid run arguments were provided")
	}

	return cfg, nil
}
