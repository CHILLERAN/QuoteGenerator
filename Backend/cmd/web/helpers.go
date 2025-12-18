package main

import (
	"net/http"
	"strings"

	"github.com/CHILLERAN/QuoteGenerator/internal/config"
)

func serverError(app *config.Application, w http.ResponseWriter, r *http.Request, err error) {
	var(
		method = r.Method
		url = r.URL
	)

	app.Logger.Info(err.Error(), "Method", method, "URL", url)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func replaceEncodingProblem(quote *string) {
	*quote = strings.ReplaceAll(*quote, "ÔÇÖ", "'")
	*quote = strings.ReplaceAll(*quote, "ÔÇô", "–")
}