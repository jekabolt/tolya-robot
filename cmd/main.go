package main

import (
	"log"
	"time"

	"github.com/caarlos0/env"
	"github.com/jekabolt/tolya-robot/bot"
)

func main() {
	log.SetFlags(log.Lshortfile)
	cfg := &bot.Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("main:env.Parse [%v]", err.Error())
	}

	b, err := cfg.Init()
	if err != nil {
		log.Fatalf("main:cfg.Init [%v]", err.Error())
	}
	err = b.SetHandlers()
	if err != nil {
		log.Fatalf("main:b.SetHandlers [%v]", err.Error())
	}

	time.Sleep(time.Minute * 100)

}
