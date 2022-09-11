package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/xo/dburl"
)

func NewConnection() (*sql.DB, error) {
	url := os.Getenv("UNIAR_DSN")
	if url == "" {
		return nil, fmt.Errorf("UNIAR_DSNが設定されていません")
	}

	db, err := dburl.Open(url)
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー")
	}

	return db, nil
}
