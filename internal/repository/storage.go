package repository

import (
	"github.com/shabkir02/go-shortener/internal/models"
	"github.com/shabkir02/go-shortener/internal/utils"
)

type Storage interface {
	GetURL(HashURL string) models.ShortURLStruct
	AddURL(u models.ShortURLStruct) (models.ShortURLStruct, error)
	GetAllURLs(userID string) []models.ShortURLStruct
}

type ShortURLs = map[string]models.ShortURLStruct

type shortURL struct {
	Storage
}

func NewStorageURL() *shortURL {
	return &shortURL{}
}

func (s *shortURL) GetURL(HashURL string) models.ShortURLStruct {
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

func (s *shortURL) AddURL(u models.ShortURLStruct) (models.ShortURLStruct, error) {
	allURLs := s.GetAllURLs("")
	var currentURLs []models.ShortURLStruct

	for _, v := range allURLs {
		if u.HashURL != v.HashURL {
			currentURLs = append(currentURLs, v)
		}
	}
	urls := append(currentURLs, u)

	cfg := utils.GetConfig()
	producer, err := utils.NewProducer(cfg.FilePatn)
	if err != nil {
		return models.ShortURLStruct{}, err
	}
	defer producer.Close()
	producer.WriteURL(&urls)

	return u, nil
}

func (s *shortURL) GetAllURLs(userID string) []models.ShortURLStruct {
	cfg := utils.GetConfig()
	consumer, err := utils.NewConsumer(cfg.FilePatn)
	if err != nil {
		return nil
	}
	defer consumer.Close()

	urls, err := consumer.ReadURLs()
	if err != nil {
		return []models.ShortURLStruct{}
	}

	if userID != "" {
		var userURLs []models.ShortURLStruct

		for _, v := range *urls {
			for _, ID := range v.UserIDs {
				if ID == userID {
					userURLs = append(userURLs, v)
					break
				}
			}
		}
		return userURLs
	}

	return *urls
}
