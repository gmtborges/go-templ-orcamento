package main

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/handlers"
	"github.com/gmtborges/orcamento-auto/middlewares"
	"github.com/gmtborges/orcamento-auto/repositories"
	"github.com/gmtborges/orcamento-auto/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)

	e := echo.New()

	sessionKey := os.Getenv("SESSION_KEY_SECRET")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionKey))))

	e.GET("/static/*", echo.WrapHandler(static()))

	indexHandler := handlers.NewIndexHandler()
	e.GET("/", indexHandler.Index)

	policyHandler := handlers.NewPolicyHandler()
	e.GET("/politica-privacidade", policyHandler.Index)

	userRepo := repositories.NewPostgresUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	loginHandler := handlers.NewLoginHandler(userSvc)
	e.GET("/login", loginHandler.Index)
	e.POST("/login", loginHandler.Create)
	e.DELETE("/logout", loginHandler.Logout)

	biddingRepo := repositories.NewPostgresBiddingRepository(db)
	biddingSvc := services.NewBiddingService(biddingRepo)
	biddingHandler := handlers.NewBiddingsHandler(biddingSvc)

	authGroup := e.Group("")
	authGroup.Use(middlewares.Authentication(userSvc))
	authGroup.GET("/orcamentos", biddingHandler.Index)
	authGroup.GET("/ofertas", biddingHandler.Index)

	e.Logger.Fatal(e.Start(":3000"))
}
