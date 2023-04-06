package main

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := OpenDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to database")

	return connection, nil
}
