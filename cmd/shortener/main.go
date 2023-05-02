package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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
		fmt.Println(r.Header.Get("Content-Type"))
		// разрешаем запросы cross-domain
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Println(string(b))

		w.Write([]byte(b))
		w.WriteHeader(http.StatusCreated)
	case "GET":
		w.Header().Set("Location", "http://localhost:8080")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	mux := http.NewServeMux()
	handler1 := http.HandlerFunc(GetHandler)
	mux.Handle("/", Conveyor(handler1, middleware))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
