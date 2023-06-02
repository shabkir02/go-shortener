package services

import (
	"errors"
	"net/http"

	"github.com/shabkir02/go-shortener/internal/models"
	"github.com/shabkir02/go-shortener/internal/repository"
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

func (h *URLService) WriteURL(hashURL string, URL string) (models.ShortURLStruct, error) {
	a := []models.ShortURLStruct{{HashURL: hashURL, URL: URL}}

	newEntry, err := h.storage.AddURL(a)
	if err != nil {
		return models.ShortURLStruct{}, errors.New("somthing went wrong")
	}

	return newEntry[0], nil
}

func (h *URLService) GetURL(hashURL string, URL string) (s models.ShortURLStruct, status int) {
	ch := hashURL
	if hashURL == "" && URL == "" {
		return models.ShortURLStruct{}, http.StatusBadRequest
	}
	if hashURL == "" && URL != "" {
		ch = h.generateHash(URL)
	}

	u := h.storage.GetURL(ch)

	if u == (models.ShortURLStruct{}) && URL == "" {
		return models.ShortURLStruct{}, http.StatusBadRequest
	}
	if u == (models.ShortURLStruct{}) && URL != "" {
		newEntry, err := h.WriteURL(ch, URL)
		if err != nil {
			return models.ShortURLStruct{}, http.StatusBadRequest
		}

		return newEntry, http.StatusCreated
	}

	return u, http.StatusOK

}
