package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/CHILLERAN/QuoteGenerator/internal/config"
	"github.com/CHILLERAN/QuoteGenerator/internal/models"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	db, err := openDB("mysql", "quotereader:password@/randomQuotesDB")

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &config.Application{
		TodoModel: &models.TodoModel{DB: db},
		Logger: logger,
	}

	logger.Info("Starting server", "Host", "http://localhost:4000")

	err = http.ListenAndServe(":4000", routes(app))

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}