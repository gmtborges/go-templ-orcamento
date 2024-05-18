package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"

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

	ep := models.Employer{
		Type: "company",
	}
	models.Employers().DeleteAll(ctx, db)
	err = ep.Insert(ctx, db, boil.Infer())
	if err != nil {
		panic(err)
	}
	cp := models.Company{
		Name:       "My assoc",
		EmployerID: ep.ID,
	}
	models.Companies().DeleteAll(ctx, db)
	err = cp.Insert(ctx, db, boil.Infer())
	if err != nil {
		panic(err)
	}
}
