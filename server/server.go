package server

import (
	"crypto/rsa"

	"github.com/jekabolt/tolya-robot/schemas"
)

type Server struct {
	DecryptKey     *rsa.PrivateKey
	DB             *schemas.DB
	Port           string
	SubmitHTMLPath string
	SubmitJSPath   string
	SubmitCSSPath  string
}
