package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jekabolt/tolya-robot/schemas"
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

func readImage(url string) (*tgbotapi.FileBytes, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("readImage:http.Get:err [%v]", err.Error())
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("readImage:ioutil.ReadAll:err [%v]", err.Error())
	}

	return &tgbotapi.FileBytes{Name: url, Bytes: bs}, nil
}

func buildPostMessage(post *schemas.Post) (*tgbotapi.PhotoConfig, error) {
	// PostMessage
	// "<b>%s</b> \n Цена: <b>%d</b>\n Размеры: <b>%v</b>\n %s \n <a href=" + "%s" + ">ссылка</a>  \n %s"

	fb, err := readImage(post.Image)
	if err != nil {
		return nil, fmt.Errorf("BuildPostMessage:readImage:err [%v]", err.Error())
	}
	msg := tgbotapi.NewPhotoUpload(132962764, *fb)
	msg.Caption = fmt.Sprintf(PostMessage, post.Title, post.Price, post.ShoeSizes, post.AboutText, post.Link, post.Hashtags)
	msg.ParseMode = "HTML"
	return &msg, nil
}
