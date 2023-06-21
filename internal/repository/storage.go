package repository

import (
	"errors"

	"github.com/shabkir02/go-shortener/internal/models"
	"github.com/shabkir02/go-shortener/internal/utils"
)

type Storage interface {
	GetURL(HashURL string) models.ShortURLStruct
	AddURL(u []models.ShortURLStruct) ([]models.ShortURLStruct, error)
	GetAllURLs() []models.ShortURLStruct
}

type ShortURLs = map[string]models.ShortURLStruct

type shortURL struct {
	urlMap ShortURLs
	Storage
}

func NewStorageURL() *shortURL {
	return &shortURL{urlMap: make(ShortURLs)}
}

func (s *shortURL) GetURL(HashURL string) models.ShortURLStruct {
	v, ok := s.urlMap[HashURL]

	if !ok {
		cfg := utils.GetConfig()
		consumer, err := utils.NewConsumer(cfg.FilePatn)
		if err != nil {
			return models.ShortURLStruct{}
		}
		defer consumer.Close()

		urls, err := consumer.ReadURLs()
		if err != nil || len(*urls) <= 0 {
			return models.ShortURLStruct{}
		}

		m := models.ShortURLStruct{}
		for _, v := range *urls {
			if v.HashURL == HashURL {
				m = v
				break
			}
		}

		return m
	}

	return v
}

func (s *shortURL) AddURL(u []models.ShortURLStruct) ([]models.ShortURLStruct, error) {
	if len(u) <= 0 {
		return []models.ShortURLStruct{}, errors.New("somthing went wrong")
	}

	for _, v := range u {
		s.urlMap[v.HashURL] = v
	}

	cfg := utils.GetConfig()
	producer, err := utils.NewProducer(cfg.FilePatn)
	if err != nil {
		return []models.ShortURLStruct{}, err
	}
	defer producer.Close()

	var mappedURLs []models.ShortURLStruct

	for k, v := range s.urlMap {
		mappedURLs = append(mappedURLs, models.ShortURLStruct{HashURL: k, URL: v.URL})
	}

	producer.WriteURL(&mappedURLs)

	return u, nil
}

func (s *shortURL) GetAllURLs() []models.ShortURLStruct {
	cfg := utils.GetConfig()
	consumer, err := utils.NewConsumer(cfg.FilePatn)
	if err != nil {
		return nil
	}
	defer consumer.Close()

	urls, err := consumer.ReadURLs()

	if err != nil || len(*urls) <= 0 {
		return []models.ShortURLStruct{}
	}

	return *urls
}
