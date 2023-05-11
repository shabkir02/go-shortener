package transport

import (
	"io"
	"net/http"

	"github.com/shabkir02/go-shortener/internal/services"
	"github.com/shabkir02/go-shortener/internal/utils"
)

type Handler struct {
	url *services.URLService
}

func NewURLHandler(u *services.URLService) *Handler {
	return &Handler{url: u}
}

func (h Handler) WriteURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}

	body, err := io.ReadAll(r.Body)
	u := string(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len([]rune(u)) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if h.url.URLMap[u] != "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(utils.GenerateURL(r.Host, h.url.URLMap[u])))
		return
	}

	newURL := h.url.WriteURL(u)
	ur := utils.GenerateURL(r.Host, newURL)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + ur))
}

func (h Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}

	if len([]rune(r.URL.Path)) > 1 {
		u := h.url.GetURL(r.URL.Path)

		if u != "" {
			w.Header().Set("Location", u)
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}

		http.Error(w, "bad request", http.StatusBadRequest)
	}

	http.Error(w, "", http.StatusBadRequest)
}
