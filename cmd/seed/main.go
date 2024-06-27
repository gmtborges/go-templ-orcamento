package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/gustavomtborges/orcamento-auto/auth"
	"github.com/gustavomtborges/orcamento-auto/db"
	"github.com/gustavomtborges/orcamento-auto/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env: %v", err)
	}
	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)
	ctx := context.Background()

	cleanUp(ctx, db)

	cp := models.Company{
		Name: "Orçamento Auto",
	}
	if err := cp.Insert(ctx, db, boil.Infer()); err != nil {
		panic(err)
	}

	hash, err := auth.GeneratePasswordHash("123")
	if err != nil {
		log.Fatal(err)
	}
	u := models.User{
		Name:      "Gustavo",
		Email:     null.StringFrom("admin@test.com"),
		Password:  hash,
		CompanyID: cp.ID,
	}
	if err := u.Insert(ctx, db, boil.Infer()); err != nil {
		panic(err)
	}
}

func cleanUp(ctx context.Context, db *sql.DB) {
	_, err := models.Companies().DeleteAll(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
}
