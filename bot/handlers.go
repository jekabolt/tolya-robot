package bot

import (
	"fmt"
	"log"

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

func (b *Bot) start(upd tgbotapi.Update) {

	link, err := b.generateLink(upd.Message.From.String(), "api/v1.0/submit")
	if err != nil {
		log.Printf("start:generateLink:err: [%s]", err.Error())
	}
	// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, link)

	b.Bot.Send(msg)
}
