package server

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
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

func (s *Server) submit(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	chatID := chi.URLParam(r, "id")

	if s.DB.IsJoined(chatID) {
		consumer, err := UnmarshalConsumer(r.Body)
		if err != nil {
			log.Printf("submit:UnmarshalConsumer: [%v]", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		s.DB.SubmitConsumer(consumer)
		if err != nil {
			log.Printf("submit:s.DB.SubmitConsumer: [%v]", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		s.DB.UpdateSubmitted(chatID)
		if err != nil {
			log.Printf("submit:s.DB.UpdateSubmitted: [%v]", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(chatID))
}

func (s *Server) submitHTML(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile(s.SubmitHTMLPath)
	if err != nil {
		log.Printf("callbackHTML:ioutil.ReadFile: [%v]", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func (s *Server) submitJS(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile(s.SubmitJSPath)
	if err != nil {
		log.Printf("callbackJS:ioutil.ReadFile: [%v]", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}
