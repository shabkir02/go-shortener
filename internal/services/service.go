package services

import (
	"fmt"
	"strings"

	hashids "github.com/speps/go-hashids/v2"
)

type URLService struct {
	URLMap map[string]string
}

func NewService() *URLService {
	return &URLService{URLMap: make(map[string]string)}
}

func (h *URLService) WriteURL(URL string) string {
	fmt.Println(h.URLMap)
	hd := hashids.NewData()
	hd.Salt = string(URL)
	hd.MinLength = 7
	hwd, _ := hashids.NewWithData(hd)
	e, _ := hwd.Encode([]int{10, 543, 321, 22})

	h.URLMap[string(URL)] = e

	fmt.Println(h.URLMap)
	return e
}

func (h *URLService) GetURL(hashURL string) string {
	sl := strings.Split(hashURL, "/")[1]
	var reqURL string

	for k, v := range h.URLMap {
		if v == sl {
			if strings.Contains(k, "https://") || strings.Contains(k, "http://") {
				reqURL = k
			} else {
				reqURL = "http://" + k
			}
		}
	}

	return reqURL

}
