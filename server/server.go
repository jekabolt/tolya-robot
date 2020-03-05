package server

import (
	"github.com/jekabolt/tolya-robot/schemas"
)

type Server struct {
	DB             *schemas.DB
	PostChan       chan *schemas.Post
	Port           string
	SubmitHTMLPath string
	SubmitJSPath   string
	SubmitCSSPath  string
}
