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

	static "github.com/gmtborges/orcamento-auto"
	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/handlers"
	"github.com/gmtborges/orcamento-auto/middlewares"
	"github.com/gmtborges/orcamento-auto/repos"
	"github.com/gmtborges/orcamento-auto/services"
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

	e.GET("/static/*", echo.WrapHandler(static.Serve()))

	indexHandler := handler.NewIndexHandler()
	e.GET("/", indexHandler.Index)

	policyHandler := handler.NewPolicyHandler()
	e.GET("/politica-privacidade", policyHandler.Index)

	userRepo := repos.NewPgUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userSvc)
	e.GET("/entrar", authHandler.Index)
	e.POST("/entrar", authHandler.Login)
	e.DELETE("/sair", authHandler.Logout)

	biddingRepo := repos.NewPgBiddingRepository(db)
	autoCategoryRepo := repos.NewPgAutoCategoryRepository(db)
	biddingSvc := services.NewBiddingService(biddingRepo, autoCategoryRepo)
	biddingHandler := handler.NewBiddingHandler(biddingSvc)

	offerRepo := repos.NewPgOfferRepository(db)
	offerSvc := services.NewOfferService(offerRepo)
	offerHandler := handler.NewOfferHandler(offerSvc, biddingSvc)

	biddingGroup := e.Group("")
	biddingGroup.Use(middlewares.Authentication(userSvc))
	biddingGroup.GET("/orcamentos", biddingHandler.Index)
	biddingGroup.GET("/orcamentos/novo", biddingHandler.New)
	biddingGroup.GET("/orcamentos/visualizar/:id", biddingHandler.Show)
	biddingGroup.GET("/orcamentos/editar/:id", biddingHandler.Edit)
	biddingGroup.POST("/orcamentos/salvar", biddingHandler.Create)
	biddingGroup.GET("/propostas/item/:id", offerHandler.GetOffersByBiddingItemID)

	offersGroup := e.Group("")
	offersGroup.Use(middlewares.Authentication(userSvc))
	offersGroup.GET("/propostas", offerHandler.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Error on starting server")
	}
}
