package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jekabolt/tolya-robot/bot"
	"github.com/jekabolt/tolya-robot/schemas"
	"github.com/jekabolt/tolya-robot/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURL       string `env:"MONGO_URL" envDefault:"mongodb://localhost:27017"`
	BotToken       string `env:"TELEGRAM_BOT_TOKEN" envDefault:""`
	BaseURL        string `env:"BASE_URL" envDefault:"https://vk.com/"`
	ServerPort     string `env:"SERVER_PORT" envDefault:"8080"`
	SubmitHTMLPath string `env:"SUBMIT_HTML_PATH" envDefault:"./web/index.html"`
	SubmitJSPath   string `env:"SUBMIT_JS_PATH" envDefault:"./web/js/index.js"`
	SubmitCSSPath  string `env:"SUBMIT_CSS_PATH" envDefault:"./web/css/index.css"`
	BotDebug       bool   `env:"BOT_DEBUG" envDefault:"true"`
}

func (c *Config) InitBot() (*bot.Bot, error) {
	b, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		return nil, fmt.Errorf("Init:NewBotAPI:err: [%s]", err.Error())
	}
	b.Debug = c.BotDebug

	db, err := c.InitDB()
	if err != nil {
		log.Fatalf("Init:bot.InitDB:err [%v]", err.Error())
	}

	return &bot.Bot{
		Bot:     b,
		BaseURL: c.BaseURL,
		DB:      db,
	}, nil
}

func (c *Config) InitServer() (*server.Server, error) {
	s := &server.Server{
		SubmitHTMLPath: c.SubmitHTMLPath,
		SubmitJSPath:   c.SubmitJSPath,
		SubmitCSSPath:  c.SubmitCSSPath,
	}

	db, err := c.InitDB()
	if err != nil {
		log.Fatalf("Init:bot.InitDB:err [%v]", err.Error())
	}
	s.DB = db
	s.Port = c.ServerPort

	return s, nil
}

func (c *Config) InitDB() (*schemas.DB, error) {
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
