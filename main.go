package main

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/handler"
	"github.com/gmtborges/orcamento-auto/middlewares"
	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/svc"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)

	e := echo.New()
	e.Use(middleware.Recover())
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
	authGroup.Use(middlewares.Authentication(userSvc))
	authGroup.GET("/orcamentos", biddingHandler.Index)
	authGroup.GET("/orcamentos/novo", biddingHandler.New)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Error on starting server")
	}
}
