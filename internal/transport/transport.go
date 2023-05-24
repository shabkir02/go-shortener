package transport

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shabkir02/go-shortener/internal/repository"
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := string(body)

	if len([]rune(u)) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, status := h.url.GetURL("", u)
	if status == http.StatusBadRequest {
		return
	}

	if status != http.StatusBadRequest {
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(status)
		w.Write([]byte(utils.GenerateURL(r.Host, c.HashURL)))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (h Handler) WhriteURLJSON(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var value ShortenURL
	json.Unmarshal(body, &value)

	v, status := h.url.GetURL("", value.URL)
	if status == http.StatusBadRequest {
		http.Error(w, "does not exist", http.StatusBadRequest)
		return
	}

	if status != http.StatusBadRequest {
		m, err := json.Marshal(ShortenURLRes{
			ResultURL: utils.GenerateURL(r.Host, v.HashURL),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(status)
		w.Write(m)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (h Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")

	u, _ := h.url.GetURL(hash, "")
	if u == (repository.ShortURLStruct{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", utils.ValidateURL(u.URL))
	w.WriteHeader(http.StatusTemporaryRedirect)
}
