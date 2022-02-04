package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Store interface {
	Open() error
	Close() error
}

var schema = `
CREATE TABLE IF NOT EXISTS table1
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	column1 TEXT,
	column2 TEXT

);
`

type dbStore struct {
	db *sqlx.DB
}

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}
