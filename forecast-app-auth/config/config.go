package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	HTTP             uint64
	DatabaseUser     string
	DatabasePassword string
	DatabaseAddress  string
	DatabaseName     string
}

func Parse() (*Config, error) {
	cfg := &Config{}

	flag.Uint64Var(&cfg.HTTP, "http", 8080, "listened by app port")
	flag.StringVar(&cfg.DatabaseAddress, "database-address", "", "address of database")
	flag.StringVar(&cfg.DatabaseName, "database-name", "", "name of database")
	flag.Parse()

	if flag.NArg() > 0 {
		return nil, fmt.Errorf("invalid run arguments were provided")
	}

	cfg.DatabaseUser = os.Getenv("DB_USER")
	cfg.DatabasePassword = os.Getenv("DB_PASSWORD")

	if len(cfg.DatabaseUser) == 0 || len(cfg.DatabasePassword) == 0 {
		return nil, fmt.Errorf("database user and password must be provided as environment variables")
	}

	return cfg, nil
}
