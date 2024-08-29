package main

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gmtborges/orcamento-auto/auth"
	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env: %v", err)
	}
	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)
	cleanUp(db)

	cp := types.Company{
		Name:      "Orçamento Auto",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tx := db.MustBegin()
	compID := int64(0)
	err = tx.QueryRow(`INSERT INTO companies 
  (name, created_at, updated_at) 
  VALUES ($1, $2, $3) RETURNING id`, cp.Name, cp.CreatedAt, cp.UpdatedAt).Scan(&compID)
	if err != nil {
		log.Fatal(err)
	}
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	hash, err := auth.GeneratePasswordHash("123")
	if err != nil {
		log.Fatal(err)
	}
	u := types.User{
		Name:      "Gustavo",
		Email:     "admin@test.com",
		Password:  hash,
		CompanyID: compID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tx = db.MustBegin()
	_, err = tx.NamedExec(`INSERT INTO users 
  (name, email, password, company_id, created_at, updated_at) 
  VALUES (:name, :email, :password, :company_id, :created_at, :updated_at)`, u)
	if err != nil {
		log.Fatal(err)
	}
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func cleanUp(db *sqlx.DB) {
	_, err := db.Exec("DELETE FROM companies")
	if err != nil {
		log.Fatal(err)
	}
}
