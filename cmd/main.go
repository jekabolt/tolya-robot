package main

import (
	"log"

	"github.com/jekabolt/tolya-robot/schemas"

	"github.com/caarlos0/env"
	"github.com/jekabolt/tolya-robot/configs"
)

func main() {
	log.SetFlags(log.Lshortfile)
	cfg := &configs.Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("main:env.Parse [%v]", err.Error())
	}

	postChan := make(chan *schemas.Post, 10)
	b, err := cfg.InitBot()
	if err != nil {
		log.Fatalf("main:cfg.InitBot [%v]", err.Error())
	}
	b.PostChan = postChan

	go b.SetHandlers()

	s, err := cfg.InitServer()
	if err != nil {
		log.Fatalf("main:cfg.InitSever [%v]", err.Error())
	}
	s.PostChan = postChan
	log.Fatalf("server.Serve():err: [%s]", s.Serve())

}
