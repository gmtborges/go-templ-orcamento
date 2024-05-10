package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/gustavomtborges/orcamento-auto/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connStr := os.Getenv("DATABASE_URL")
	dbconn(connStr)

	e := echo.New()
	e.GET("/static/*", echo.WrapHandler(static()))
	e.GET("/", handlers.GetIndex)
	e.GET("/politica-privacidade", handlers.GetPolicy)
	e.GET("/login", handlers.GetLogin)

	e.Logger.Fatal(e.Start(":3000"))
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
