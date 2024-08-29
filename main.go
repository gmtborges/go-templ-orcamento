package main

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/handler"
	"github.com/gmtborges/orcamento-auto/middleware"
	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/svc"
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

	indexHandler := handler.NewIndexHandler()
	e.GET("/", indexHandler.Index)

	policyHandler := handler.NewPolicyHandler()
	e.GET("/politica-privacidade", policyHandler.Index)

	userRepo := repo.NewPostgresUserRepository(db)
	userSvc := svc.NewUserService(userRepo)
	loginHandler := handler.NewLoginHandler(userSvc)
	e.GET("/login", loginHandler.Index)
	e.POST("/login", loginHandler.Create)
	e.DELETE("/logout", loginHandler.Logout)

	biddingRepo := repo.NewPostgresBiddingRepository(db)
	biddingSvc := svc.NewBiddingService(biddingRepo)
	biddingHandler := handler.NewBiddingsHandler(biddingSvc)

	authGroup := e.Group("")
	authGroup.Use(middleware.Authentication(userSvc))
	authGroup.GET("/orcamentos", biddingHandler.Index)
	authGroup.GET("/ofertas", biddingHandler.Index)

	e.Logger.Fatal(e.Start(":3000"))
}
