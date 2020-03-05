package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) Serve() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.HandleFunc("/", s.healthCheck)
	r.Options("/*", handleOptions)

	r.Route("/static", func(r chi.Router) {
		r.Get("/submit/{id}", s.submitHTML)
		r.Get("/submit.js", s.submitJS)
		r.Get("/submit.css", s.submitCSS)
	})

	r.Route("/api/v1.0", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Post("/submit/{id}", s.submit)
		r.Post("/send", s.send)
	})

	log.Println("Listening on :" + s.Port)
	return http.ListenAndServe(":"+s.Port, r)

}
