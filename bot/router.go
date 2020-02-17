package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) SetHandlers() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.Bot.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("SetHandlers:GetUpdatesChan: err: [%s]", err.Error())
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		b.HandleCommand(update)
	}

	return nil
}

func (b *Bot) HandleCommand(upd tgbotapi.Update) {
	method, ok := fetchCommand(upd.Message.Text)

	if ok {
		switch method {
		case "/start":
			b.start(upd)
		case "/info":

		}
	}
}
