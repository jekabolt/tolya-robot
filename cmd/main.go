package main

import (
	"log"

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

	b, err := cfg.InitBot()
	if err != nil {
		log.Fatalf("main:cfg.InitBot [%v]", err.Error())
	}

	go b.SetHandlers()

	s, err := cfg.InitServer()
	if err != nil {
		log.Fatalf("main:cfg.InitSever [%v]", err.Error())
	}
	log.Fatalf("server.Serve():err: [%s]", s.Serve())

}
