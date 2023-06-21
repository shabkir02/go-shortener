package middleware

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func generateRandom(size int) ([]byte, error) {
	// генерируем случайную последовательность байт
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateHash() ([]byte, error) {
	src, err := generateRandom(64)
	if err != nil {
		return []byte{}, err
	}
	// создаём случайный ключ
	key, err := generateRandom(16)
	if err != nil {
		return []byte{}, err
	}
	// подписываем алгоритмом HMAC, используя SHA256
	h := hmac.New(sha256.New, key)
	h.Write(src)
	dst := h.Sum(nil)

	return dst, nil
}

func CheckUserCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := r.Cookie("user")
		if err != nil {
			h, err := generateHash()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// r.Context()
			cookie := http.Cookie{Name: "user", Value: hex.EncodeToString(h)}
			http.SetCookie(w, &cookie)
		}

		next.ServeHTTP(w, r)
	})
}
