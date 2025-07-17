package config

import "github.com/caarlos0/env/v11"

type Config struct {
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
}

func Load() (*Config, error) {
	config, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}
	return &config, nil
}
