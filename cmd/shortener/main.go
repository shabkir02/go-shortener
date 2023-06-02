package main

import (
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/shabkir02/go-shortener/internal/app"
	"github.com/shabkir02/go-shortener/internal/utils"
)

func main() {

	utils.InitConfig()
	cfg := utils.GetConfig()

	r := app.NewRouter()
	color.Green("Server started.")
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, r))
}
