package cfg

type Database struct {
	User           string `env:"DATABASE_USER"`
	Password       string `env:"DATABASE_PASSWORD"`
	Host           string `env:"DATABASE_HOST"`
	Port           int    `env:"DATABASE_PORT"`
	Name           string `env:"DATABASE_NAME"`
	SSLMode        string `env:"DATABASE_SSL_MODE" default:"disable"`
	MigrationsPath string `env:"DATABASE_MIGRATIONS_PATH"`
}

