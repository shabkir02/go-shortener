package transport

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	defer r.Body.Close()
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

	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	if h.url.URLMap[u] != "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(utils.GenerateURL(r.Host, h.url.URLMap[u])))
		return
	}

	newURL := h.url.WriteURL(u)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(utils.GenerateURL(r.Host, newURL)))
}

func (h Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	hashID := chi.URLParam(r, "id")

	u := h.url.GetURL(hashID)

	fmt.Println(u)

	if u != "" {
		w.Header().Set("Location", u)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "bad request", http.StatusBadRequest)

}
