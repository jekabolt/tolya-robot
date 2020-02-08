package bot

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

var bitSize = 2048
var tgID = "132962764"
var pubPemPath = "../certs/pub.pem"
var privPemPath = "../certs/priv.pem"

func TestGenCerts(t *testing.T) {
	priv, pub := GenerateRsaKeyPair()

	err := savePublicPEMKey(pubPemPath, *pub)
	if err != nil {
		log.Fatalf("savePublicPEMKey:err [%v]", err.Error())
	}

	err = savePEMKey(privPemPath, priv)
	if err != nil {
		log.Fatalf("savePEMKey:err [%v]", err.Error())
	}
}

func TestSign(t *testing.T) {
	pubBytes, err := ReadFile(pubPemPath)
	if err != nil {
		log.Fatalf("ReadFile(pubPemPath):err [%v]", err.Error())
	}

	privBytes, err := ReadFile(privPemPath)
	if err != nil {
		log.Fatalf("ReadFile(privPemPath):err [%v]", err.Error())
	}

	pubPem, err := ParseRsaPublicKeyFromPem(pubBytes)
	if err != nil {
		log.Fatalf("BytesToPublicKey:err [%v]", err.Error())
	}

	privPem, err := ParseRsaPrivateKeyFromPem(privBytes)
	if err != nil {
		log.Fatalf("BytesToPrivateKey:err [%v]", err.Error())
	}

	ciphertext, err := EncryptWithPublicKey([]byte(tgID), pubPem)
	if err != nil {
		log.Fatalf("EncryptWithPublicKey:err [%v]", err.Error())
	}

	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("%s", ciphertextBase64)

	decr, err := DecryptWithPrivateKey(UnBase64(ciphertextBase64), privPem)
	if err != nil {
		log.Fatalf("DecryptWithPrivateKey:err [%v]", err.Error())
	}

	if tgID == string(decr) {
		fmt.Printf("\n\n\n\ncooollllll\n\n\n")
	}

}
