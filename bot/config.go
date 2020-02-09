package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	Port         string `env:"SERVER_PORT" envDefault:"8080"`
	BotToken     string `env:"TELEGRAM_BOT_TOKEN" envDefault:""`
	SignCertPath string `env:"SIGN_SERT_PATH" envDefault:"certs/pub.pem"`
	BaseURL      string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	Debug        bool   `env:"DEBUG" envDefault:"true"`
}

func (c *Config) Init() (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		return nil, fmt.Errorf("Init:NewBotAPI:err: [%s]", err.Error())
	}
	bot.Debug = c.Debug

	pubBytes, err := ReadFile(c.SignCertPath)
	if err != nil {
		log.Fatalf("Init:bot.ReadFile:err [%v]", err.Error())
	}

	pubPem, err := ParseRsaPublicKeyFromPem(pubBytes)
	if err != nil {
		log.Fatalf("Init:bot.ParseRsaPublicKeyFromPem:err [%v]", err.Error())
	}

	return &Bot{
		Bot:     bot,
		SigKey:  pubPem,
		BaseURL: c.BaseURL,
	}, nil
}
