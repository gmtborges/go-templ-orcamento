package main

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gustavomtborges/orcamento-auto/db"
	"github.com/gustavomtborges/orcamento-auto/handlers"
	"github.com/gustavomtborges/orcamento-auto/middlewares"
	"github.com/gustavomtborges/orcamento-auto/services"
	"github.com/gustavomtborges/orcamento-auto/stores"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)

	e := echo.New()
	e.Use(middleware.Logger())

	sessionKey := os.Getenv("SESSION_KEY_SECRET")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionKey))))

	e.GET("/static/*", echo.WrapHandler(static()))

	indexHandler := handlers.NewIndexHandler()
	e.GET("/", indexHandler.Show)

	policyHandler := handlers.NewPolicyHandler()
	e.GET("/politica-privacidade", policyHandler.Show)

	userStore := stores.NewPostgresUserStore(db)
	loginSvc := services.NewLoginService(userStore)
	authSvc := services.NewAuthService()
	loginHandler := handlers.NewLoginHandler(loginSvc, authSvc)
	e.GET("/login", loginHandler.Show)
	e.POST("/login", loginHandler.Create)

	dashSvc := services.NewDashService(userStore)
	dashHandler := handlers.NewDashHandler(dashSvc)
	e.GET("/dashboard", dashHandler.Show, middlewares.Authentication)

	e.Logger.Fatal(e.Start(":3000"))
}
