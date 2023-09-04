package cfg

type Server struct {
	Host             string `env:"SERVER_HOST"`
	Port             int    `env:"SERVER_PORT"`
	Secret           string `env:"SERVER_SECRET"`
	CasbinAuthModel  string `env:"SERVER_CASBIN_AUTH_MODEL"`
	CasbinAuthPolicy string `env:"SERVER_CASBIN_AUTH_POLICY"`
	UIHost           string `env:"SERVER_UI_HOST"`
	AllowedHosts     string `env:"SERVER_ALLOWED_HOSTS" envDefault:"*"`
}
