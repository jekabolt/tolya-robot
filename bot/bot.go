package bot

import (
	"crypto/rsa"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	Bot     *tgbotapi.BotAPI
	SigKey  *rsa.PublicKey
	BaseURL string
}
