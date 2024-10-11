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

	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/services"
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

	companies := []types.Company{
		{
			Name:      "Org",
			Type:      types.CompanyTypeOrg,
			State:     "GO",
			City:      "Goiania",
			Telephone: "62998765432",
		},
		{
			Name:      "Auto Service",
			Type:      types.CompanyTypeAuto,
			State:     "GO",
			City:      "Goiania",
			Telephone: "62998765432",
		},
		{
			Name:      "Auto Product",
			Type:      types.CompanyTypeAuto,
			State:     "GO",
			City:      "Goiania",
			Telephone: "62998765432",
		},
	}

	tx := db.MustBegin()
	for i, c := range companies {
		var id int64
		err := tx.QueryRow(`INSERT INTO companies 
    (name, type, state, city, telephone, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5, now(), now()) 
    RETURNING id;
    `, c.Name, c.Type, c.State, c.City, c.Telephone).Scan(&id)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to insert into companies")
		}
		companies[i].ID = id
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	hash, err := services.GeneratePasswordHash("123")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate password hash")
	}
	users := []types.User{
		{
			Name:      "User Org",
			Email:     "org@test.com",
			Password:  hash,
			CompanyID: sql.NullInt64{Int64: companies[0].ID, Valid: true},
		},
		{
			Name:      "User Auto Product",
			Email:     "auto1@test.com",
			Password:  hash,
			CompanyID: sql.NullInt64{Int64: companies[1].ID, Valid: true},
		},
		{
			Name:      "User Auto Service",
			Email:     "auto2@test.com",
			Password:  hash,
			CompanyID: sql.NullInt64{Int64: companies[2].ID, Valid: true},
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
	      VALUES ($1, $2, $3, $4, now(), now()) 
        RETURNING id;
        `, u.Name, u.Email, u.Password, u.CompanyID).Scan(&userID)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to insert into users")
		}
		users[i].ID = userID
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	setUserRoles(db, users)
	autoCategoryIDs := seedAutoCategories(db)
	autoCategoriesCompanies := map[int64][]int64{
		companies[1].ID: autoCategoryIDs[:3],
		companies[2].ID: autoCategoryIDs[3:],
	}
	joinAutoCategoriesCompanies(db, autoCategoriesCompanies)
	seedBiddings(db, companies[0].ID, autoCategoryIDs)
}

func setUserRoles(db *sqlx.DB, users []types.User) {
	var roleID int64
	err := db.Get(&roleID, `SELECT id FROM roles WHERE name = 'ADMIN' LIMIT 1;`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get roles")
	}

	tx := db.MustBegin()
	for _, u := range users {
		_, err = tx.Exec(`INSERT INTO users_roles (user_id, role_id, created_at, updated_at) 
    VALUES ($1, $2, now(), now());`, u.ID, roleID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert into users_roles")
		}
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}
}

func seedBiddings(db *sqlx.DB, orgID int64, categoryIDs []int64) {
	var uID int64
	err := db.Get(&uID, `SELECT id FROM users WHERE company_id = $1;`, orgID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get user")
	}
	for i := 1; i <= 34; i++ {

		b := types.Bidding{
			CompanyID:    orgID,
			UserID:       uID,
			CustomerName: fmt.Sprintf("Associado %d", i),
			VehicleBrand: "Ford",
			VehicleName:  fmt.Sprintf("veiculo %d", i),
			VehicleYear:  1990 + i,
			VehicleColor: "Preto",
			Notes:        "veiculo todo fudido",
			Status:       getRandomStatus(),
			CreatedAt:    time.Now().AddDate(0, 0, -i),
		}

		bi := []types.BiddingItem{
			{
				AutoCategoryID: categoryIDs[0],
				Notes:          "",
				Status:         types.BiddingItemStatusOpen,
			},
			{
				AutoCategoryID: categoryIDs[1],
				Notes:          "Uma observacao bem grande que passa de 30 caracteres.",
				Status:         types.BiddingItemStatusOfferAccepted,
			},
			{
				AutoCategoryID: categoryIDs[2],
				Notes:          "Uma observacao bem grande que passa de 30 caracteres.",
				Status:         types.BiddingItemStatusOfferReceived,
			},
			{
				AutoCategoryID: categoryIDs[3],
				Notes:          "",
				Status:         types.BiddingItemStatusOpen,
			},
			{
				AutoCategoryID: categoryIDs[4],
				Notes:          "",
				Status:         types.BiddingItemStatusCanceled,
			},
			{
				AutoCategoryID: categoryIDs[5],
				Notes:          "",
				Status:         types.BiddingItemStatusOpen,
			},
		}

		tx := db.MustBegin()
		var biddingID int64
		err := tx.QueryRow(`
		INSERT INTO biddings (company_id, user_id, customer_name, vehicle_brand, vehicle_name, 
    vehicle_year, vehicle_color, notes, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now())
		RETURNING id;`, b.CompanyID, b.UserID, b.CustomerName, b.VehicleBrand, b.VehicleName, b.VehicleYear,
			b.VehicleColor, b.Notes, b.Status, b.CreatedAt).Scan(&biddingID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert into biddings")
		}

		for _, item := range bi {
			item.BiddingID = biddingID
			_, err := tx.Exec(`
			INSERT INTO bidding_items (bidding_id, auto_category_id, notes, status)
			VALUES ($1, $2, $3, $4);`, item.BiddingID, item.AutoCategoryID, item.Notes, item.Status)
			if err != nil {
				tx.Rollback()
				log.Fatal().Err(err).Msg("Failed to insert into bidding_items")
			}
		}

		if err := tx.Commit(); err != nil {
			log.Fatal().Err(err).Msg("Failed to commit transaction")
		}
	}
}

func getRandomStatus() types.BiddingStatus {
	status := []types.BiddingStatus{
		types.BiddingStatusAwaitingOffer,
		types.BiddingStatusPending,
		types.BiddingStatusFinished,
		types.BiddingStatusCanceled,
	}
	return status[rand.Intn(len(status))]
}

func seedAutoCategories(db *sqlx.DB) []int64 {
	autocategories := []types.AutoCategory{
		{Description: "Lanternagem", Type: types.AutoCategoryTypeService},
		{Description: "Pintura", Type: types.AutoCategoryTypeService},
		{Description: "Mecânica", Type: types.AutoCategoryTypeService},
		{Description: "Parachoque", Type: types.AutoCategoryTypeProduct},
		{Description: "Retrovisor", Type: types.AutoCategoryTypeProduct},
		{Description: "Porta", Type: types.AutoCategoryTypeProduct},
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start transaction")
	}
	var autoCategoryIDs []int64

	for _, ac := range autocategories {
		var autoCategoryID int64
		err = tx.QueryRow(`
			INSERT INTO auto_categories (description, type)
			VALUES ($1, $2)
			RETURNING id;
      `, ac.Description, ac.Type).Scan(&autoCategoryID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert into auto_categories")
		}

		autoCategoryIDs = append(autoCategoryIDs, autoCategoryID)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	return autoCategoryIDs
}

func joinAutoCategoriesCompanies(db *sqlx.DB, autoCategoriesCompanies map[int64][]int64) {
	tx, err := db.Beginx()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start transaction for companies_auto_categories")
	}
	for k, acIDs := range autoCategoriesCompanies {
		for _, acID := range acIDs {
			tx.Exec(`INSERT INTO companies_auto_categories (company_id, auto_category_id) 
      VALUES ($1, $2);`, k, acID)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit companies_auto_categories")
	}
}

func cleanUp(db *sqlx.DB) {
	_, err := db.Exec(`
		DELETE FROM users_roles;
		DELETE FROM users;
    DELETE FROM companies_auto_categories;
		DELETE FROM companies;
		DELETE FROM auto_categories;
		DELETE FROM biddings;
	`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to clean up database")
	}
}
