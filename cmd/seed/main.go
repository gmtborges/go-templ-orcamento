package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env: %v", err)
	}
	connStr := os.Getenv("DB_URL")
	db := dbConn(connStr)
	sql, err := os.ReadFile("seed.sql")
	if err != nil {
		log.Fatalf("Error on reading seed: %v", err)
	}
	seed(db, string(sql))
}

func seed(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func dbConn(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening db: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging db: %v", err)
	}
	return db
}
