package services

import (
	"strings"

	hashids "github.com/speps/go-hashids/v2"
)

type URLService struct {
	urlMap map[string]string
}

func NewService() *URLService {
	return &URLService{urlMap: make(map[string]string)}
}

func (h *URLService) CheckURL(URL string) *string {

	if v, ok := h.urlMap[URL]; ok {
		return &v
	}

	return nil
}

func (h *URLService) AddURL(key string, value string) {
	h.urlMap[key] = value
}

func (h *URLService) WriteURL(URL string) string {
	hd := hashids.NewData()
	hd.Salt = string(URL)
	hd.MinLength = 7
	hwd, _ := hashids.NewWithData(hd)
	e, _ := hwd.Encode([]int{10, 543, 321, 22})

	h.urlMap[string(URL)] = e

	return e
}

func (h *URLService) GetURL(hashURL string) string {
	var reqURL string

	for k, v := range h.urlMap {
		if v == hashURL {
			if strings.Contains(k, "https://") || strings.Contains(k, "http://") {
				reqURL = k
			} else {
				reqURL = "http://" + k
			}
		}
	}

	return reqURL

}
