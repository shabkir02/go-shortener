package app

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/shabkir02/go-shortener/internal/middleware"
	"github.com/shabkir02/go-shortener/internal/services"
	"github.com/shabkir02/go-shortener/internal/transport"
)

// var defaultCompressibleContentTypes = []string{
// 	"text/html",
// 	"text/css",
// 	"text/plain",
// 	"text/javascript",
// 	"application/javascript",
// 	"application/x-javascript",
// 	"application/json",
// 	"application/atom+xml",
// 	"application/rss+xml",
// 	"image/svg+xml",
// }

func NewRouter() chi.Router {
	service := services.NewService()
	handlers := transport.NewURLHandler(service)
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.GzipHandle)

	r.Post("/", handlers.WriteURL)
	r.Get("/{hash}", handlers.GetURL)
	r.Post("/api/shorten", handlers.WhriteURLJSON)

	return r
}
