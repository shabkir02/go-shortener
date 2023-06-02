package utils

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
	FilePatn      string `env:"FILE_STORAGE_PATH"`
}

var cfg *Config

func InitConfig() {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	f := &Config{}
	flag.StringVar(&f.ServerAddress, "a", ":8080", "Server address")
	flag.StringVar(&f.BaseURL, "b", "http://localhost:8080", "Base URL")
	flag.StringVar(&f.FilePatn, "f", "urls.json", "File Path")
	flag.Parse()

	if c.ServerAddress == "" {
		c.ServerAddress = f.ServerAddress
	}
	if c.BaseURL == "" {
		c.BaseURL = f.BaseURL
	}
	if c.FilePatn == "" {
		c.FilePatn = f.FilePatn
	}

	cfg = c
}

func GetConfig() Config {
	if cfg == nil {
		InitConfig()
	}
	return *cfg
}
