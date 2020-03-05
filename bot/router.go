package bot

import (
	"fmt"

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
	// PostMessage
	// "<b>%s</b> \n Цена: <b>%d</b>\n Размеры: <b>%v</b>\n %s \n <a href=" + "%s" + ">ссылка</a>  \n %s"

	// msg := tgbotapi.NewPhotoUpload(132962764, "/Users/jekabolt/Documents/grb-logo.jpg")
	msg := tgbotapi.NewMediaGroup(132962764, []interface{}{"/Users/jekabolt/Documents/grb-logo.jpg", "/Users/jekabolt/Documents/grb-logo.jpg", "/Users/jekabolt/Documents/grb-logo.jpg"})
	// msg.Caption = fmt.Sprintf(PostMessage, post.Title, post.Price, post.ShoeSizes, post.AboutText, post.Link, post.Hashtags)
	// msg.ParseMode = "HTML"
	// msg.ReplyMarkup
	b.Bot.Send(msg)

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
