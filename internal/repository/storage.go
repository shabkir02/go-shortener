package repository

import (
	"errors"
)

type Storage interface {
	GetURL(HashURL string) ShortURLStruct
	AddURL(u ShortURLStruct) (ShortURLStruct, error)
}

type ShortURLStruct struct {
	HashURL string
	URL     string
}
type ShortURLs = map[string]ShortURLStruct

type shortURL struct {
	urlMap ShortURLs
}

func NewStorageURL() *shortURL {
	return &shortURL{urlMap: make(ShortURLs)}
}

func (s *shortURL) GetURL(HashURL string) ShortURLStruct {
	v, ok := s.urlMap[HashURL]

	if !ok {
		return ShortURLStruct{}
	}

	return v
}

func (s *shortURL) AddURL(u ShortURLStruct) (ShortURLStruct, error) {
	s.urlMap[u.HashURL] = u

	_, ok := s.urlMap[u.HashURL]

	if !ok {
		return ShortURLStruct{}, errors.New("somthing went wrong")
	}

	return u, nil
}
