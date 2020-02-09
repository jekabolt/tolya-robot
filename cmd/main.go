package main

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/jekabolt/tolya-robot/bot"
	"github.com/jekabolt/tolya-robot/server"
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

	go b.SetHandlers()

	server := &server.Server{}
	err = env.Parse(server)
	if err != nil {
		log.Fatalf("main:env.Parse [%v]", err.Error())
	}
	err = server.Init()
	if err != nil {
		log.Fatalf("main:b.SetHandlers [%v]", err.Error())
	}
	log.Fatalf("server.Serve():err: [%s]", server.Serve())

	// time.Sleep(time.Minute * 100)

}
