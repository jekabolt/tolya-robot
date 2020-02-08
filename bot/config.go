package bot

type Config struct {
	Port     string `env:"SERVER_PORT" envDefault:"8080"`
	BotToken string `env:"TELEGRAM_BOT_TOKEN" envDefault:"1054316886:AAGzfrilNW-5HnfwrShP0VCUKx-dE2UPekM"`
	Debug    bool   `env:"DEBUG" envDefault:"true"`
}
