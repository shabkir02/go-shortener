package utils

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FilePatn      string `env:"FILE_STORAGE_PATH" envDefault:"urls.json"`
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

	if c == (&Config{}) {
		cfg = f
		return
	}

	cfg = c
}

func GetConfig() Config {
	if cfg == nil {
		InitConfig()
	}
	return *cfg
}
