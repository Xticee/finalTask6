package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type app struct {
	Server *http.Server
	Logger *log.Logger
}

func CreateApp(logger *log.Logger) *app {
	return &app{
		Server: newServer(logger),
		Logger: logger,
	}
}

func newServer(logger *log.Logger) *http.Server {
	router := chi.NewRouter()
	router.Get("/", handlers.GetForm)
	router.Post("/upload", handlers.UploadForm)

	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	return &server
}
