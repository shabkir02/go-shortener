package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/shabkir02/go-shortener/internal/middleware"
	"github.com/shabkir02/go-shortener/internal/models"
	"github.com/shabkir02/go-shortener/internal/repository"
	"github.com/shabkir02/go-shortener/internal/utils"
	hashids "github.com/speps/go-hashids/v2"
)

type URLService struct {
	storage repository.Storage
}

func NewService() *URLService {
	return &URLService{storage: repository.NewStorageURL()}
}

func (h *URLService) generateHash(URL string) string {
	hd := hashids.NewData()
	hd.Salt = string(URL)
	hd.MinLength = 7
	hwd, _ := hashids.NewWithData(hd)

	e, _ := hwd.Encode([]int{10, 543, 321, 22})

	return e
}

func (h *URLService) WriteURL(u models.ShortURLStruct) (models.ShortURLStruct, error) {
	newEntry, err := h.storage.AddURL(u)
	if err != nil {
		return models.ShortURLStruct{}, errors.New("somthing went wrong")
	}

	return newEntry, nil
}

func (h *URLService) GetURL(m models.ShortURLStruct, ctx context.Context) (s models.ShortURLStruct, status int) {
	userID := ctx.Value(middleware.UserIDContextKey).(string)

	ch := m.HashURL
	if m.IsEmpty() {
		return models.ShortURLStruct{}, http.StatusBadRequest
	}
	if m.HashURL == "" && m.URL != "" {
		ch = h.generateHash(m.URL)
	}

	u := h.storage.GetURL(ch)

	if u.IsEmpty() && m.URL == "" {
		return models.ShortURLStruct{}, http.StatusBadRequest
	}
	if u.IsEmpty() && m.URL != "" {
		newEntry, err := h.WriteURL(models.ShortURLStruct{HashURL: ch, URL: m.URL, UserIDs: []string{userID}})
		if err != nil {
			return models.ShortURLStruct{}, http.StatusBadRequest
		}

		return newEntry, http.StatusCreated
	}
	if !u.UserExist(userID) {
		var ms models.ShortURLStruct

		if u.IsEmpty() {
			ms = models.ShortURLStruct{HashURL: ch, URL: m.URL, UserIDs: []string{userID}}
		} else {
			userIDs := append(u.UserIDs, userID)
			ms = models.ShortURLStruct{HashURL: u.HashURL, URL: u.URL, UserIDs: userIDs}
		}

		newEntry, err := h.WriteURL(ms)
		if err != nil {
			return models.ShortURLStruct{}, http.StatusBadRequest
		}

		return newEntry, http.StatusOK
	}

	return u, http.StatusOK
}

func (h *URLService) GetAllURLs(ctx context.Context) []models.AllURLsStruct {
	userID := ctx.Value(middleware.UserIDContextKey).(string)

	m := h.storage.GetAllURLs(userID)
	e := make([]models.AllURLsStruct, len(m))
	cfg := utils.GetConfig()

	for i, v := range m {
		e[i] = models.AllURLsStruct{
			OriginalURL: v.URL,
			ShortURL:    utils.GenerateURL(cfg.BaseURL, v.HashURL),
		}
	}

	return e
}
