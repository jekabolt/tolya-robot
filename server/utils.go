package server

import (
	"fmt"
	"net/url"

	"github.com/jekabolt/tolya-robot/bot"
)

func (s *Server) ciphertextDecode(ciphertextBase64UrlEncoded string) ([]byte, error) {
	ciphertextBase64, err := url.PathUnescape(ciphertextBase64UrlEncoded)
	if err != nil {
		return nil, fmt.Errorf("ciphertextDecode:url.PathUnescape:err [%v]", err.Error())
	}
	ciphertext, err := bot.UnBase64(ciphertextBase64)
	if err != nil {
		return nil, fmt.Errorf("ciphertextDecode:url.UnBase64:err [%v]", err.Error())
	}
	id, err := bot.DecryptWithPrivateKey(ciphertext, s.DecryptKey)
	if err != nil {
		return nil, fmt.Errorf("ciphertextDecode:url.PathUnescape:err [%v]", err.Error())
	}
	return id, err
}
