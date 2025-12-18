package models

import (
	"database/sql"
	"fmt"
	"math/rand"
)

type Todo struct {
	ID     int
	Author string
	Quote  string
}

type TodoModel struct {
	DB *sql.DB
}

func (tm *TodoModel) GetRandomQuote() (Todo, error) {
	var amountOfQuotes int
	var randomTodo Todo

	countQuery := `SELECT COUNT(*)
	FROM QUOTES`

	err := tm.DB.QueryRow(countQuery).Scan(&amountOfQuotes)

	if err != nil {
		return Todo{}, err
	}

	randomId := rand.Int31n(int32(amountOfQuotes)) + 1

	query := `SELECT *
	FROM QUOTES
	WHERE ID = ?`

	err = tm.DB.QueryRow(query, randomId).Scan(&randomTodo.ID, &randomTodo.Author, &randomTodo.Quote)

	if err != nil {
		return Todo{}, err
	}

	return randomTodo, nil
}

func (tm *TodoModel) GetQuoteByWord(word string) ([]Todo, error) {
	wordParameter := fmt.Sprintf("%v%v%v", "%", word, "%")

	query := `SELECT *
	FROM QUOTES
	WHERE LOWER(QUOTE) LIKE LOWER(?)`

	rows, err := tm.DB.Query(query, wordParameter) 

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var quotes []Todo

	for rows.Next() {
		var quote Todo

		err := rows.Scan(&quote.ID, &quote.Author, &quote.Quote)

		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}  

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}