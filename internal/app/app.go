package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shabkir02/go-shortener/internal/services"
	"github.com/shabkir02/go-shortener/internal/transport"
)

func NewRouter() chi.Router {
	service := services.NewService()
	handlers := transport.NewURLHandler(service)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{id}", handlers.GetURL)
	r.Post("/", handlers.WriteURL)

	return r
}
