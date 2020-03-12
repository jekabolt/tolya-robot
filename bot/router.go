package bot

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jekabolt/tolya-robot/schemas"
)

func (b *Bot) SetHandlers() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.Bot.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("SetHandlers:GetUpdatesChan: err: [%s]", err.Error())
	}

	go func() {
		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}
			b.HandleCommand(update)
		}
	}()

	go func() {
		for post := range b.PostChan {
			fmt.Println("post ", post)
			if post == nil { // ignore any non-Message Updates
				continue
			}
			b.HandlePost(post)
		}

	}()

	return nil
}

func (b *Bot) HandlePost(post *schemas.Post) {

	ids, err := b.DB.FetchConsumersForPost(post)
	if err != nil {
		log.Printf("HandlePost:b.DB.FetchConsumersForPost:err: [%v]", err.Error())
		return
	}

	msg, err := buildPostMessage(post)
	if err != nil {
		log.Printf("HandlePost:buildPostMessage:err: [%v]", err.Error())
		return
	}

	// https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once according with
	batchSize := 30
	batchTime := time.Second * 1
	c := 1
	for _, id := range ids {
		c++
		if batchSize < c {
			time.Sleep(batchTime)
		}
		chatID, _ := strconv.Atoi(id)
		msg.ChatID = int64(chatID)
		_, err = b.Bot.Send(msg)
		if err != nil {
			log.Printf("HandlePost:b.Bot.Send:err: [%v]", err.Error())
		}
	}

}

func (b *Bot) HandleCommand(upd tgbotapi.Update) {
	method, ok := fetchCommand(upd.Message.Text)

	if ok {
		//commands
		switch method {
		case "/start":
			b.start(upd)
		}
	} else {
		switch upd.Message.Text {
		case "FAQ":
			b.handleFAQ(upd)
		case "Настройки":
			b.handleSettings(upd)
		case "Лучшие предложения":
			b.handleBestOffers(upd)
		}
	}

}
