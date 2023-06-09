package transport

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/shabkir02/go-shortener/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_WriteURL(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		urlRes      string
	}
	tests := []struct {
		name    string
		want    want
		request string
		urlBody string
	}{
		{
			name: "Сохранение записи",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusCreated,
				urlRes:      "http://example.com/g8SrEcqnUX",
			},
			request: "/",
			urlBody: "https://music.yandex.ru/artist/8095900",
		},
		{
			name: "Сохранение записи",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusOK,
				urlRes:      "http://example.com/g8SrEcqnUX",
			},
			request: "/",
			urlBody: "https://music.yandex.ru/artist/8095900",
		},
		{
			name: "Сохранение записи",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusCreated,
				urlRes:      "http://example.com/gLSwmULGCx",
			},
			request: "/",
			urlBody: "https://pkg.go.dev/net/http",
		},
	}

	service := services.NewService()
	handlers := NewURLHandler(service)
	r := chi.NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader(tt.urlBody))
			w := httptest.NewRecorder()

			hFun := http.HandlerFunc(handlers.WriteURL)
			hFun(w, request)

			result := w.Result()
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := io.ReadAll(result.Body)
			require.NoError(t, err)

			err = result.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, tt.want.urlRes, string(userResult))
		})
	}

}

func TestHandler_GetURL(t *testing.T) {
	type want struct {
		statusCode int
		urlRes     string
	}
	tests := []struct {
		name    string
		want    want
		request string
	}{
		{
			name: "Чтение записи(нет)",
			want: want{
				statusCode: http.StatusBadRequest,
				urlRes:     "",
			},
			request: "/5nj4k35",
		},
		{
			name: "Чтение записи(нет)",
			want: want{
				statusCode: http.StatusBadRequest,
				urlRes:     "",
			},
			request: "/",
		},
	}

	service := services.NewService()
	handlers := NewURLHandler(service)
	r := chi.NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	handlers.url.URLMap["https://music.yandex.ru/artist/8095900"] = "g8SrEcqnUX"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			w := httptest.NewRecorder()

			hFun := http.HandlerFunc(handlers.GetURL)
			hFun(w, request)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.urlRes, string(result.Header.Get("Location")))
		})
	}

}
