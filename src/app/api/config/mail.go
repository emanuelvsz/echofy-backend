package cfg

type Mail struct {
	SMTPHost     string `env:"MAIL_SMTP_HOST"`
	SMTPPort     string `env:"MAIL_SMTP_PORT"`
	FromName     string `env:"MAIL_FROM_NAME"`
	FromAddress  string `env:"MAIL_FROM_ADDRESS"`
	FromPassword string `env:"MAIL_FROM_PASSWORD"`
}