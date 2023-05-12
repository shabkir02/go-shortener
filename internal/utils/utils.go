package utils

import (
	"strings"
)

func GenerateURL(host string, path string) string {
	var sb strings.Builder
	sb.WriteString(host)
	sb.WriteString("/")
	sb.WriteString(path)
	s := sb.String()

	if strings.Contains(s, "https://") || strings.Contains(s, "http://") {
		return s
	} else {
		return "http://" + s
	}
}

// type Middleware func(http.Handler) http.Handler

// func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
// 	for _, middleware := range middlewares {
// 		h = middleware(h)
// 	}
// 	return h
// }

// func middleware(next http.Handler) http.Handler {
// 	// собираем Handler приведением типа
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// разрешаем запросы cross-domain
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		next.ServeHTTP(w, r)
// 	})
// }
