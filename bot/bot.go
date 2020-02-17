package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jekabolt/tolya-robot/schemas"
)

type Bot struct {
	Bot     *tgbotapi.BotAPI
	DB      *schemas.DB
	BaseURL string
}
