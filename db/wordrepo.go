package db

import (
	"database/sql"
	"fmt"
)

// SQLWordRepo finds words if they exist
type SQLWordRepo struct {
	ConnString string
}

// Search finds a word if it exists
func (wr SQLWordRepo) Search(prefix string, limit int) ([]string, error) {
	sql, err := sql.Open("mysql", wr.ConnString)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT word FROM entries WHERE word LIKE '%v%%'", prefix)
	if limit != -1 {
		query += fmt.Sprintf(" LIMIT %v", limit)
	}

	rows, err := sql.Query(query)

	if err != nil {
		return nil, err
	}

	var words []string
	for rows.Next() {
		var word string
		_ = rows.Scan(&word)
		words = append(words, word)
	}

	return words, nil
}
