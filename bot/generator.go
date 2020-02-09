package bot

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func (b *Bot) generateLink(id, method string) (string, error) {
	ciphertext, err := EncryptWithPublicKey([]byte(id), b.SigKey)
	if err != nil {
		return "", fmt.Errorf(" EncryptWithPublicKey:err [%v]", err.Error())
	}
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	return b.BaseURL + "/" + method + "/" + url.PathEscape(ciphertextBase64), err
}
