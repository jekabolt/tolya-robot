package bot

import (
	"strconv"

	"github.com/jekabolt/tolya-robot/schemas"

	"github.com/aws/aws-sdk-go/aws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) start(upd tgbotapi.Update) {

	b.DB.InitialSubmit(&schemas.TGUser{
		User:      upd.Message.From,
		ChatID:    upd.Message.Chat.ID,
		Submitted: false,
	})

	link := b.BaseURL + "api/v1.0/submit/" + strconv.Itoa(int(upd.Message.Chat.ID))

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, link)

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
