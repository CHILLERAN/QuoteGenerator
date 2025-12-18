package config

import (
	"log/slog"

	"github.com/CHILLERAN/QuoteGenerator/internal/models"
)

type Application struct {
	TodoModel *models.TodoModel
	Logger *slog.Logger
}