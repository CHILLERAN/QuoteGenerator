package main

import (
	"net/http"

	"github.com/CHILLERAN/QuoteGenerator/internal/config"
)

func routes(app *config.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", home(app))
	mux.Handle("GET /{word}", getQuoteWithWord(app))

	return mux
}