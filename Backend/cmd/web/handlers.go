package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/CHILLERAN/QuoteGenerator/internal/config"
)

type quoteHandler func(w http.ResponseWriter, r *http.Request)

func (qh quoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qh(w, r)
} 

func home(app *config.Application) quoteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		quote, err := app.TodoModel.GetRandomQuote()
		
		if err != nil {
			serverError(app, w, r, err)
			return 
		}

		w.WriteHeader(http.StatusOK)
		
		replaceEncodingProblem(&quote.Quote)			
		fmt.Fprintf(w ,"%v\n\n-%v-", quote.Quote, quote.Author)
	}
}

func getQuoteWithWord(app *config.Application) quoteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		quotes, err := app.TodoModel.GetQuoteByWord(r.PathValue("word"))
		
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			app.Logger.Error(err.Error())
			http.NotFound(w, r)

			return 
		}else if (err != nil) {
			serverError(app, w, r, err)

			return 
		}

		w.WriteHeader(http.StatusOK)
		
		for _, quote := range quotes {
			replaceEncodingProblem(&quote.Quote)			
			fmt.Fprintf(w ,"%v\n\n-%v-\n\n\n", quote.Quote, quote.Author)
		}

	}
}
