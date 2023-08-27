package env

type DB struct {
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port     int64  `env:"DB_PORT" envDefault:"5432"`
	Database string `env:"DB_DATABASE" envDefault:"kikumemodb"`
}
