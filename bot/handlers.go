package bot

import (
	"fmt"
	"strconv"

	"github.com/jekabolt/tolya-robot/schemas"

	"github.com/aws/aws-sdk-go/aws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) start(upd tgbotapi.Update) {

	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("FAQ"),
			tgbotapi.NewKeyboardButton("Настройки"),
			tgbotapi.NewKeyboardButton("Лучшие предложения"),
		),
	)

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)
	msg.ReplyMarkup = numericKeyboard

	m, err := b.Bot.Send(msg)
	fmt.Println("m ", m)
	fmt.Println("err ", err)

	b.DB.InitialSubmit(&schemas.TGUser{
		User:      upd.Message.From,
		ChatID:    upd.Message.Chat.ID,
		Submitted: false,
	})

	link := b.BaseURL + "static/submit/" + strconv.Itoa(int(upd.Message.Chat.ID))

	msg = tgbotapi.NewMessage(upd.Message.Chat.ID, link)

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{Text: "Выбрать одежду",
					URL: aws.String(link),
				},
			},
		},
	}
	msg.Text = startMessage

	b.Bot.Send(msg)
}
