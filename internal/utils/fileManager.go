package utils

import (
	"encoding/json"
	"os"

	"github.com/shabkir02/go-shortener/internal/models"
)

type producer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewProducer(fileName string) (*producer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return nil, err
	}
	return &producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}
func (p *producer) WriteURL(event *[]models.ShortURLStruct) error {
	return p.encoder.Encode(&event)
}

func (p *producer) Close() error {
	return p.file.Close()
}

type consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func NewConsumer(fileName string) (*consumer, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (c *consumer) ReadURLs() (*[]models.ShortURLStruct, error) {
	event := []models.ShortURLStruct{}
	if err := c.decoder.Decode(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (c *consumer) Close() error {
	return c.file.Close()
}
