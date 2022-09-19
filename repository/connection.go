package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./uniar.db")
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー")
	}

	return db, nil
}
