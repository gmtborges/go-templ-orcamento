package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/db"
	"github.com/gustavomtborges/orcamento-auto/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)

	e := echo.New()
	e.GET("/static/*", echo.WrapHandler(static()))
	e.GET("/", handlers.ShowIndex)
	e.GET("/politica-privacidade", handlers.ShowPolicy)
	e.GET("/login", handlers.ShowLogin)
	e.POST("/login", func(c echo.Context) error {
		return handlers.PostLogin(c, db)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
