package bot

import (
	"fmt"
	"strconv"

	"github.com/jekabolt/tolya-robot/schemas"

	"github.com/aws/aws-sdk-go/aws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("FAQ"),
		tgbotapi.NewKeyboardButton("Настройки"),
		tgbotapi.NewKeyboardButton("Лучшие предложения"),
	),
)

var backKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

func (b *Bot) start(upd tgbotapi.Update) {

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)
	msg.ReplyMarkup = mainKeyboard

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

func (b *Bot) handleFAQ(upd tgbotapi.Update) {

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, FAQMessage)

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{Text: "Ссылка на статью",
					URL: aws.String("https://telegra.ph/Gajd-Voprosy-i-otvety-02-26"),
				},
			},
		},
	}

	b.Bot.Send(msg)
}

func (b *Bot) handleSettings(upd tgbotapi.Update) {

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, settingsMessage)

	link := b.BaseURL + "static/submit/" + strconv.Itoa(int(upd.Message.Chat.ID))
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{Text: "Изменить настройки",
					URL: aws.String(link),
				},
			},
		},
	}

	_, _ = b.Bot.Send(msg)
}

func (b *Bot) handleBestOffers(upd tgbotapi.Update) {
	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, BestOffers)
	msg.ParseMode = "HTML"
	_, err := b.Bot.Send(msg)
	fmt.Println("\n\n\n\n handleSettings:", err)
}
