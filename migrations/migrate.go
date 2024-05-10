package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connStr := os.Getenv("DATABASE_URL")
	db := dbconn(connStr)
	migrations(db)

}

//go:embed *.sql
var embedMigrations embed.FS

func migrations(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	fmt.Println("running migrations...")
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}

func dbconn(connStr string) *sql.DB {
	if connStr == "" {
		log.Fatalf("DATABASE_URL environment variable not set")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	return db
}
