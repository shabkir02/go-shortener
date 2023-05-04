package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	hashids "github.com/speps/go-hashids/v2"
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

func generateUrl(host string, path string) string {
	var sb strings.Builder
	sb.WriteString(host)
	sb.WriteString("/")
	sb.WriteString(path)

	return sb.String()
}

var m = make(map[string]string)

func GetHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		b, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if len([]rune(string(b))) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if m[string(b)] != "" {
			w.WriteHeader(http.StatusOK)

			w.Write([]byte(generateUrl(r.Host, m[string(b)])))
			return
		}

		hd := hashids.NewData()
		hd.Salt = string(b)
		hd.MinLength = 7
		h, _ := hashids.NewWithData(hd)
		e, _ := h.Encode([]int{10, 543, 321, 22})

		m[string(b)] = e

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(generateUrl(r.Host, e)))
	case "GET":
		if len([]rune(r.URL.Path)) > 1 {
			sl := strings.Split(r.URL.Path, "/")[1]
			var reqUrl string

			for k, v := range m {
				if v == sl {
					reqUrl = k
				}
			}

			if reqUrl != "" {
				w.Header().Set("Location", reqUrl)
				w.WriteHeader(http.StatusTemporaryRedirect)
				return
			}

		}

		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	mux := http.NewServeMux()
	handler1 := http.HandlerFunc(GetHandler)
	mux.Handle("/", Conveyor(handler1, middleware))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
