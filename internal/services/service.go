package services

import (
	"net/http"

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

func (h *URLService) GetURL(hashURL string, URL string) (s repository.ShortURLStruct, status int) {
	ch := hashURL
	if hashURL == "" && URL == "" {
		return repository.ShortURLStruct{}, http.StatusBadRequest
	}
	if hashURL == "" && URL != "" {
		ch = h.generateHash(URL)
	}

	u := h.storage.GetURL(ch)

	if u == (repository.ShortURLStruct{}) && URL == "" {
		return repository.ShortURLStruct{}, http.StatusBadRequest
	}
	if u == (repository.ShortURLStruct{}) && URL != "" {
		newEntry, err := h.storage.AddURL(repository.ShortURLStruct{HashURL: ch, URL: URL})
		if err != nil {
			return repository.ShortURLStruct{}, http.StatusBadRequest
		}

		return newEntry, http.StatusCreated
	}

	return u, http.StatusOK

}
