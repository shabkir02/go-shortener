package app

import (
	"log"
	"net/http"

	"github.com/shabkir02/go-shortener/internal/services"
	"github.com/shabkir02/go-shortener/internal/transport"
)

type Middleware func(http.Handler) http.Handler

func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func middleware(next http.Handler) http.Handler {
	// собираем Handler приведением типа
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// разрешаем запросы cross-domain
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	mux := http.NewServeMux()
	service := services.NewService()
	handlers := transport.NewURLHandler(service)

	mux.Handle("/", Conveyor(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url(w, r, handlers)
	}), middleware))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func url(w http.ResponseWriter, r *http.Request, handlers *transport.Handler) {

	switch r.Method {
	case http.MethodGet:
		handlers.GetURL(w, r)
	case http.MethodPost:
		handlers.WriteURL(w, r)
	default:
		http.Error(w, "method not allowed", 500)
	}

}
