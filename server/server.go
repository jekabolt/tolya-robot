package server

import (
	"crypto/rsa"
	"log"

	"github.com/jekabolt/tolya-robot/bot"
)

type Server struct {
	Port            string `env:"SERVER_PORT" envDefault:"8080"`
	DecryptCertPath string `env:"DECRYPT_CERT_PATH" envDefault:"certs/priv.pem"`
	DecryptKey      *rsa.PrivateKey
}

func (s *Server) Init() error {
	privBytes, err := bot.ReadFile(s.DecryptCertPath)
	if err != nil {
		log.Fatalf("Init:bot.ReadFile:err [%v]", err.Error())
	}

	privPem, err := bot.ParseRsaPrivateKeyFromPem(privBytes)
	if err != nil {
		log.Fatalf("Init:bot.ParseRsaPrivateKeyFromPem:err [%v]", err.Error())
	}
	s.DecryptKey = privPem

	return nil
}
