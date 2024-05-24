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
	connStr := os.Getenv("DB_URL")
	db := dbconn(connStr)
	if len(os.Args) < 2 {
		fmt.Print("Usage: go run ./cmd/migrate <up|down|status>")
		os.Exit(0)
	}
	cmd := os.Args[1]
	migrations(db, cmd)

}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrations(db *sql.DB, cmd string) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	switch cmd {
	case "up":
		fmt.Println("running migrations...")
		if err := goose.Up(db, "migrations"); err != nil {
			panic(err)
		}
	case "down":
		fmt.Println("rollback migrations...")
		if err := goose.Down(db, "migrations"); err != nil {
			panic(err)
		}
	}
}

func dbconn(connStr string) *sql.DB {
	if connStr == "" {
		log.Fatalf("DB_URL environment variable not set")
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
