package utils

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	Port     string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL  string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FilePatn string `env:"FILE_STORAGE_PATH" envDefault:"urls.json"`
}

var cfg *Config

func InitConfig() {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	cfg = c
}

func GetConfig() Config {
	return *cfg
}
