package bot

import (
	"encoding/json"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func fetchCommand(msg string) (string, bool) {
	if len(msg) == 0 {
		return "", false
	}
	if msg[0] == '/' {
		ss := strings.Split(msg, " ")
		return ss[0], true
	}
	return "", false
}

func TgUserToJson(user *tgbotapi.User) ([]byte, error) {
	return json.Marshal(user)
}
