package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	Port     string `env:"SERVER_PORT" envDefault:"8080"`
	BotToken string `env:"TELEGRAM_BOT_TOKEN" envDefault:""`
	Debug    bool   `env:"DEBUG" envDefault:"true"`
}

func (c *Config) Init() (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		return nil, fmt.Errorf("Init:NewBotAPI:err: [%s]", err.Error())
	}
	bot.Debug = c.Debug

	return &Bot{
		Bot: bot,
	}, nil
}
