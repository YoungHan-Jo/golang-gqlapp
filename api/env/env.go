package env

import (
	env_v9 "github.com/caarlos0/env/v9"
)

type DB struct {
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port     int64  `env:"DB_PORT" envDefault:"5432"`
	Database string `env:"DB_DATABASE" envDefault:"kikumemodb"`
}

type Server struct {
	Port int64 `env:"SERVER_PORT" envDefault:"8090"`
}

type Env struct {
	DB
	AllowOrigin string `env:"ALLOW_ORIGIN" envDefault:"http://api.local.gqlpracapp.com"`
	Server
}

var e Env

func init() {
	if err := env_v9.Parse(&e); err != nil {
		panic(err)
	}
}

func Get() *Env {
	return &e
}
