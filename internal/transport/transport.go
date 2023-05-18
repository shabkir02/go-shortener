package transport

import (
	"encoding/json"
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
type ShortenURL struct {
	URL string `json:"url"`
}
type ShortenURLRes struct {
	ResultURL string `json:"result"`
}

func NewURLHandler(u *services.URLService) *Handler {
	return &Handler{url: u}
}

func (h Handler) WriteURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	u := string(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len([]rune(u)) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "text/plain; charset=utf-8")

	c := h.url.CheckURL(u)

	if c != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(utils.GenerateURL(r.Host, *c)))
		return
	}

	newURL := h.url.WriteURL(u)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(utils.GenerateURL(r.Host, newURL)))
}

func (h Handler) WhriteUrlJSON(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var value ShortenURL
	json.Unmarshal(body, &value)

	w.Header().Set("Content-type", "application/json; charset=utf-8")

	v := h.url.CheckURL(value.URL)

	if v == nil {
		newURL := h.url.WriteURL(value.URL)

		m, err := json.Marshal(ShortenURLRes{
			ResultURL: utils.GenerateURL(r.Host, newURL),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(m)
		return
	}

	m, err := json.Marshal(ShortenURLRes{
		ResultURL: utils.GenerateURL(r.Host, *v),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(m)
}

func (h Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")

	u := h.url.GetURL(hash)

	fmt.Println(u)

	if u != "" {
		w.Header().Set("Location", u)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "bad request", http.StatusBadRequest)

}
