package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/auth"
	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error on loading .env: %v")
	}
	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)
	cleanUp(db)

	org := types.Company{
		Name:     "Org 1",
		Type:     types.CompanyTypeOrg,
		Province: "GO",
		City:     "Goiania",
	}

	tx := db.MustBegin()
	orgID := int64(0)
	err = tx.QueryRow(`INSERT INTO companies 
  (name, type, province, city, created_at, updated_at) 
  VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`, org.Name, org.Type, org.Province, org.City).Scan(&orgID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to insert company")
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	auto := types.Company{
		Name:     "Auto 1",
		Type:     types.CompanyTypeAuto,
		Province: "GO",
		City:     "Goiania",
	}

	tx = db.MustBegin()
	autoID := int64(0)
	err = tx.QueryRow(`INSERT INTO companies 
  (name, type, province, city, created_at, updated_at) 
  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, auto.Name, auto.Type, auto.Province, auto.City, auto.CreatedAt, auto.UpdatedAt).Scan(&autoID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to insert company")
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	hash, err := auth.GeneratePasswordHash("123")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate password hash")
	}
	users := []types.User{
		{
			Name:      "User Org",
			Email:     "org@test.com",
			Password:  hash,
			CompanyID: sql.NullInt64{Int64: orgID, Valid: true},
		},
		{
			Name:      "User Auto",
			Email:     "auto@test.com",
			Password:  hash,
			CompanyID: sql.NullInt64{Int64: autoID, Valid: true},
		},
		{
			Name:      "User Standalone",
			Email:     "stand@test.com",
			CompanyID: sql.NullInt64{Int64: 0, Valid: false},
			Password:  hash,
		},
	}

	tx = db.MustBegin()
	for i, u := range users {
		var userID int64
		err = tx.QueryRow(`INSERT INTO users 
	      (name, email, password, company_id, created_at, updated_at) 
	      VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`, u.Name, u.Email, u.Password, u.CompanyID).Scan(&userID)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to insert user")
		}
		users[i].ID = userID
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	setUserRoles(db, users)
	autoCategoryIDs := seedAutoCategories(db)
	seedBidding(db, orgID, autoCategoryIDs)
}

func setUserRoles(db *sqlx.DB, users []types.User) {
	var roleID int64
	err := db.Get(&roleID, `SELECT id FROM roles WHERE name = 'ADMIN'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get role ID")
	}

	tx := db.MustBegin()
	for _, u := range users {
		_, err = tx.Exec(`INSERT INTO users_roles (user_id, role_id, created_at, updated_at) VALUES ($1, $2, now(), now())`,
			u.ID, roleID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert user role")
		}
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}
}

func seedBidding(db *sqlx.DB, orgID int64, categoryIDs []int64) {
	var userID int64
	err := db.Get(&userID, `SELECT id FROM users WHERE company_id = $1`, orgID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get user")
	}
	for i := 1; i <= 34; i++ {

		b := types.Bidding{
			CompanyID:    orgID,
			UserID:       userID,
			CustomerName: fmt.Sprintf("Customer %d", i),
			VehicleBrand: "Ford",
			VehicleName:  fmt.Sprintf("Vehicle %d", i),
			VehicleYear:  1990 + i,
			VehicleColor: "Black",
			Status:       getRandomStatus(),
			CreatedAt:    time.Now().AddDate(0, 0, -i),
		}

		bi := []types.BiddingItem{
			{Description: sql.NullString{String: fmt.Sprintf("Description %d", i)}, AutoCategoryID: categoryIDs[0]},
			{Description: sql.NullString{String: fmt.Sprintf("Description %d", i)}, AutoCategoryID: categoryIDs[1]},
			{Description: sql.NullString{String: fmt.Sprintf("Description %d", i)}, AutoCategoryID: categoryIDs[2]},
			{Description: sql.NullString{String: fmt.Sprintf("Description %d", i)}, AutoCategoryID: categoryIDs[3]},
			{Description: sql.NullString{String: fmt.Sprintf("Description %d", i)}, AutoCategoryID: categoryIDs[4]},
			{Description: sql.NullString{String: fmt.Sprintf("Description %d", i)}, AutoCategoryID: categoryIDs[5]},
		}

		tx := db.MustBegin()
		var biddingID int64
		err := tx.QueryRow(`
		INSERT INTO biddings (company_id, user_id, customer_name, vehicle_brand, vehicle_name, vehicle_year, vehicle_color, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`, b.CompanyID, b.UserID, b.CustomerName, b.VehicleBrand, b.VehicleName, b.VehicleYear, b.VehicleColor, b.Status, b.CreatedAt).Scan(&biddingID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert bidding")
		}

		for _, item := range bi {
			item.BiddingID = biddingID
			_, err := tx.Exec(`
			INSERT INTO bidding_items (bidding_id, description, auto_category_id)
			VALUES ($1, $2, $3)`, item.BiddingID, item.Description, item.AutoCategoryID)
			if err != nil {
				tx.Rollback()
				log.Fatal().Err(err).Msg("Failed to insert bidding item")
			}
		}

		if err := tx.Commit(); err != nil {
			log.Fatal().Err(err).Msg("Failed to commit transaction")
		}
	}
}

func getRandomStatus() types.BiddingStatus {
	statuses := []types.BiddingStatus{
		types.BiddingStatusAwaitingOffers,
		types.BiddingStatusPending,
		types.BiddingStatusFinished,
		types.BiddingStatusCanceled,
	}
	return statuses[rand.Intn(len(statuses))]
}

func seedAutoCategories(db *sqlx.DB) []int64 {
	autoCategories := []types.AutoCategory{
		{Name: "Lanternagem", Type: types.CategoryTypeService},
		{Name: "Pintura", Type: types.CategoryTypeService},
		{Name: "Mecânica", Type: types.CategoryTypeService},
		{Name: "Parachoque", Type: types.CategoryTypeProduct},
		{Name: "Retrovisor", Type: types.CategoryTypeProduct},
		{Name: "Porta", Type: types.CategoryTypeProduct},
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start transaction")
	}
	var autoCategoryIDs []int64

	for _, c := range autoCategories {
		var autoCategoryID int64
		err = tx.QueryRow(`
			INSERT INTO auto_categories (name, type)
			VALUES ($1, $2)
			RETURNING id`, c.Name, c.Type).Scan(&autoCategoryID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert category")
		}

		autoCategoryIDs = append(autoCategoryIDs, autoCategoryID)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	return autoCategoryIDs
}

func cleanUp(db *sqlx.DB) {
	_, err := db.Exec(`
		DELETE FROM users_roles;
		DELETE FROM users;
		DELETE FROM companies;
		DELETE FROM auto_categories;
		DELETE FROM biddings;
	`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to clean up database")
	}
}
