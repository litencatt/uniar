package repository

import (
	"database/sql"
	"fmt"
	"os/user"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection() (*sql.DB, error) {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	uniarPath := user.HomeDir + "/.uniar"
	dbPath := uniarPath + "/uniar.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー")
	}

	return db, nil
}
