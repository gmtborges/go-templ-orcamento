package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Conn(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}
	return db
}
