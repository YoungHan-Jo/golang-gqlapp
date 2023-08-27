package server

import (
	"database/sql"
	"time"
)

func newDB(dbSource string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
