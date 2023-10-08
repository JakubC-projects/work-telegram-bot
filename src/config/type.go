package config

type Config struct {
	Server struct {
		Port        int    `envconfig:"PORT"`
		CertFile    string `envconfig:"CERT_FILE"`
		CertKeyFile string `envconfig:"CERT_KEY_FILE"`
	}
	TelegramApiKey string `envconfig:"TELEGRAM_API_KEY"`
}
