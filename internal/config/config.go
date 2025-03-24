package config

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppEnv string `env:"APP_ENV,default=local"`

	HTTPPort string `env:"HTTP_PORT,default=:8080"`

	Postgres PostgresConfig

	Jwt JWTConfig
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST,default=localhost"`
	Port     string `env:"POSTGRES_PORT,default=5432"`
	User     string `env:"POSTGRES_USER,default=postgres"`
	Password string `env:"POSTGRES_PASSWORD,default=postgres"`
	DBName   string `env:"POSTGRES_DB,default=movie_db"`
	SSLMode  string `env:"POSTGRES_SSLMODE,default=disable"`
}

type JWTConfig struct {
	SecretKey       string `env:"JWT_SECRET"`
	AccessTokenTTL  int    `env:"JWT_ACCESS_TTL_MIN,default=15"`
	RefreshTokenTTL int    `env:"JWT_REFRESH_TTL_HOURS,default=720"`
}

func (p PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode,
	)
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		return cfg, fmt.Errorf("failed to load config: %w", err)
	}
	return cfg, nil
}
