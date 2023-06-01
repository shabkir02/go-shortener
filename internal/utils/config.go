package utils

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FilePatn      string `env:"FILE_STORAGE_PATH" envDefault:"urls.json"`
}

var cfg *Config

func InitConfig() {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	if c == (&Config{}) {
		flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "Server address")
		flag.StringVar(&c.BaseURL, "b", c.BaseURL, "Base URL")
		flag.StringVar(&c.FilePatn, "f", c.FilePatn, "File Path")
		flag.Parse()
	}

	cfg = c
}

func GetConfig() Config {
	return *cfg
}
