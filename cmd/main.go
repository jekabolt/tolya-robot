package main

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/jekabolt/tolya-robot/bot"
)

func main() {
	log.SetFlags(log.Lshortfile)
	cfg := &bot.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("main:env.Parse [%v]", err.Error())
	}

}
