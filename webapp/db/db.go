package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1/alura_store?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
