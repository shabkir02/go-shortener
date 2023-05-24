package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env/v8"
	"github.com/fatih/color"
	"github.com/shabkir02/go-shortener/internal/app"
)

type config struct {
	Port string `env:"SERVER_ADDRESS" envDefault:":8080"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	r := app.NewRouter()
	color.Green("Server started.")
	log.Fatal(http.ListenAndServe(cfg.Port, r))
}
