package main

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/shabkir02/go-shortener/internal/app"
)

func main() {
	r := app.NewRouter()
	color.Green("Server started.")
	http.ListenAndServe(":8080", r)
}
