package config

import (
	"github.com/caarlos0/env"
)

// Config all app variables are stored here
type Config struct {
	// Postgres config
	Username  string `env:"DB_USERNAME,required" envDefault:"postgres"`
	Password  string `env:"DB_PASSWORD,required" envDefault:"postgres"`
	Host      string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port      int    `env:"DB_PORT" envDefault:"5432"`
	DbName    string `env:"DB_NAME,required"`
	SslEnable bool   `env:"DB_SSL_ENABLE" envDefault:"false"`

	// App config
	LogLevel   string `env:"LOG_LEVEL" envDefault:"debug"`
	ServerPort string `env:"SERVER_PORT" envDefault:"8081"`
}

// New returns a new Config struct
func New() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
