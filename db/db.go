package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Conn(connStr string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}
	return db
}
