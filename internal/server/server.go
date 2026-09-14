package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	r := chi.NewRouter()

	r.Get("/", handlers.HandleIndex)
	r.Post("/upload", handlers.HandleUpload)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 5,
	}

	return &Server{
		Logger: logger,
		Server: s,
	}
}
