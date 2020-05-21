package configs

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jekabolt/tolya-robot/bot"
	"github.com/jekabolt/tolya-robot/schemas"
	"github.com/jekabolt/tolya-robot/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURL        string `env:"MONGO_URL" envDefault:"mongodb://localhost:27017"`
	BotToken        string `env:"TELEGRAM_BOT_TOKEN" envDefault:""`
	BaseURL         string `env:"BASE_URL" envDefault:"http://dotmarket.me/"`
	ServerPort      string `env:"SERVER_PORT" envDefault:"8080"`
	SubmitHTMLPath  string `env:"SUBMIT_HTML_PATH" envDefault:"./web/index.html"`
	SuccessHTMLPath string `env:"SUBMIT_HTML_PATH" envDefault:"./web/success.html"`
	SubmitJSPath    string `env:"SUBMIT_JS_PATH" envDefault:"./web/js/index.js"`
	SubmitCSSPath   string `env:"SUBMIT_CSS_PATH" envDefault:"./web/css/index.css"`
	DBPassword      string `env:"DB_PASSWORD" envDefault:"kek"`
	BotDebug        bool   `env:"BOT_DEBUG" envDefault:"true"`
}

func (c *Config) InitBot(db *schemas.DB, postChan chan *schemas.Post) (*bot.Bot, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	b, err := tgbotapi.NewBotAPIWithClient(c.BotToken, client)
	if err != nil {
		return nil, fmt.Errorf("Init:NewBotAPI:err: [%s]", err.Error())
	}
	b.Debug = c.BotDebug

	return &bot.Bot{
		Bot:      b,
		BaseURL:  c.BaseURL,
		DB:       db,
		PostChan: postChan,
	}, nil
}

func (c *Config) InitServer(db *schemas.DB, postChan chan *schemas.Post) *server.Server {
	s := &server.Server{
		SubmitHTMLPath:  c.SubmitHTMLPath,
		SubmitJSPath:    c.SubmitJSPath,
		SubmitCSSPath:   c.SubmitCSSPath,
		SuccessHTMLPath: c.SuccessHTMLPath,
		Port:            c.ServerPort,
		DB:              db,
		PostChan:        postChan,
	}
	return s
}

func (c *Config) InitDB() (*schemas.DB, error) {
	log.Printf("mongo url: %s", c.MongoURL)
	var db = &schemas.DB{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(c.MongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Init:mongo.Connect:err [%v]", err.Error())
	}
	db.Client = client
	db.ConsumersCollection = client.Database(schemas.DBName).Collection(schemas.ConsumersCollectionName)
	db.JoinedCollection = client.Database(schemas.DBName).Collection(schemas.JoinedCollectionName)

	return db, err
}
