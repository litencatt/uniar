package repository

import (
	"database/sql"
	"fmt"
	"os"
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
	if p, ok := os.LookupEnv("UNIAR_DB_PATH"); ok {
		dbPath = p
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー")
	}

	return db, nil
}
