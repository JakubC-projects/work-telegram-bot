package config

type Config struct {
	Server struct {
		Port        int    `envconfig:"PORT"`
		CertFile    string `envconfig:"CERT_FILE"`
		CertKeyFile string `envconfig:"CERT_KEY_FILE"`
	}
	TelegramApiKey string `envconfig:"TELEGRAM_API_KEY"`
	WorkApi        struct {
		ResultsUrl string `envconfig:"WORK_API_RESULTS_URL"`
		FormUrl    string `envconfig:"WORK_API_FORM_URL"`
	}
	Redis struct {
		Url      string `envconfig:"REDIS_URL"`
		Password string `envconfig:"REDIS_PASSWORD"`
	}
}
