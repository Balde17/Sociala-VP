package queryBuilder

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB(dataSrc string) (*DB, error) {
	_, err := os.Stat(dataSrc)
	if err == nil {
		db, err := sql.Open("sqlite3", dataSrc)
		if err != nil {
			return nil, err
		}
		if err = db.Ping(); err != nil {
			db.Close()
			return nil, err
		}

		return &DB{DB: db}, nil
	} else {
		return nil, err
	}
}

func (db *DB) Close() error {
	return db.DB.Close()
}
