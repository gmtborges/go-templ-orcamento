package main

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/db"
	"github.com/gustavomtborges/orcamento-auto/handlers"
	"github.com/gustavomtborges/orcamento-auto/middlewares"
	"github.com/gustavomtborges/orcamento-auto/repositories"
	"github.com/gustavomtborges/orcamento-auto/services"
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
	authSvc := services.NewAuthService(userRepo)
	loginHandler := handlers.NewLoginHandler(authSvc)
	e.GET("/login", loginHandler.Index)
	e.POST("/login", loginHandler.Create)
	e.DELETE("/logout", loginHandler.Logout)

	biddingRepo := repositories.NewPostgresBiddingRepository(db)
	biddingSvc := services.NewBiddingService(biddingRepo)
	biddingHandler := handlers.NewBiddingsHandler(biddingSvc)
	// e.GET("/orcamentos", biddingHandler.Index, middlewares.Authentication)
	compGroup := e.Group("")
	compGroup.Use(middlewares.Authentication)
	compGroup.Use(middlewares.Authentication)
	e.GET("/orcamentos", biddingHandler.Index)
	e.GET("/ofertas", biddingHandler.Index)

	e.Logger.Fatal(e.Start(":3000"))
}
