package config

import "os"

type Config struct {
	Host string
	Port string
}

func Load() *Config {
	cfg := &Config{
		Host: "0.0.0.0",
		Port: "8080",
	}

	if port := os.Getenv("DOCKARR_PORT"); port != "" {
		cfg.Port = port
	}

	if host := os.Getenv("DOCKARR_HOST"); host != "" {
		cfg.Host = host
	}

	return cfg
}
