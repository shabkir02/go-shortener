package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shabkir02/go-shortener/internal/services"
	"github.com/shabkir02/go-shortener/internal/transport"
	"github.com/shabkir02/go-shortener/internal/utils"
)

func NewRouter() chi.Router {
	service := services.NewService()
	handlers := transport.NewURLHandler(service)
	r := chi.NewRouter()

	urls := utils.InitFileManager()

	for _, v := range *urls {
		service.WriteURL(v.HashURL, v.URL)
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{hash}", handlers.GetURL)
	r.Post("/", handlers.WriteURL)
	r.Post("/api/shorten", handlers.WhriteURLJSON)

	return r
}
