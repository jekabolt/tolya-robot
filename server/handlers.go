package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jekabolt/tolya-robot/schemas"
)

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) seen(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	ciphertextBase64UrlEncoded := chi.URLParam(r, "id")

	id, err := s.ciphertextDecode(ciphertextBase64UrlEncoded)
	if err != nil {
		log.Printf("seen:ciphertextDecode:err [%v]", err.Error())
	}

	//TODO: mongo first in

	w.WriteHeader(http.StatusOK)
	w.Write(id)
}

func (s *Server) submit(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	ciphertextBase64UrlEncoded := chi.URLParam(r, "id")

	id, err := s.ciphertextDecode(ciphertextBase64UrlEncoded)
	if err != nil {
		log.Printf("submit:ciphertextDecode:err [%v]", err.Error())
	}

	u := &schemas.User{}
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("submit:ioutil.ReadAll:err [%v]", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rawBody, u)
	if err != nil {
		log.Printf("generatePaymentURL:json.Unmarshal: [%v]", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Println(u)
	//TODO: mongo set all info

	w.WriteHeader(http.StatusOK)
	w.Write(id)
}
