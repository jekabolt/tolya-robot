package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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
