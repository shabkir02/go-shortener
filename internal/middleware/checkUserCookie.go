package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gofrs/uuid"
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
	src, err := uuid.NewV4()
	if err != nil {
		return []byte{}, err
	}
	// создаём случайный ключ
	key, err := uuid.NewV4()
	if err != nil {
		return []byte{}, err
	}
	// подписываем алгоритмом HMAC, используя SHA256
	h := hmac.New(sha256.New, []byte(key.String()))
	h.Write([]byte(src.String()))
	dst := h.Sum(nil)

	return dst, nil
}

type contextKey string

const (
	UserIDContextKey contextKey = "user"
)

func CheckUserCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("user")
		var userHash string

		if err != nil {
			h, err := generateHash()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			he := hex.EncodeToString(h)
			cookie := http.Cookie{Name: "user", Value: he}
			http.SetCookie(w, &cookie)

			userHash = he
		} else {
			userHash = c.Value
		}

		r.Context()
		ctx := context.WithValue(r.Context(), UserIDContextKey, userHash)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
