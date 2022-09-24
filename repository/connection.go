package repository

import (
	"database/sql"
	"fmt"
)

func NewConnection(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー")
	}

	return db, nil
}
