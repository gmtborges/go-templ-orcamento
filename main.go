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

	politicaPrivacidadeHandler := handler.NewPoliticaPrivacidadeHandler()
	e.GET("/politica-privacidade", politicaPrivacidadeHandler.Index)

	usuarioRepo := repo.NewPostgresUsuarioRepository(db)
	usuarioSvc := svc.NewUsuarioService(usuarioRepo)
	autenticacaoHandler := handler.NewAutenticacaoHandler(usuarioSvc)
	e.GET("/entrar", autenticacaoHandler.Index)
	e.POST("/entrar", autenticacaoHandler.Login)
	e.DELETE("/sair", autenticacaoHandler.Logout)

	orcamentoRepo := repo.NewPgOrcamentoRepository(db)
	autoCategoriaRepo := repo.NewPgAutoCategoriaRepository(db)
	orcamentoSvc := svc.NewOrcamentoService(orcamentoRepo, autoCategoriaRepo)
	orcamentoHandler := handler.NewOrcamentoHandler(orcamentoSvc)

	propostaRepo := repo.NewPgPropostaRepository(db)
	propostaSvc := svc.NewPropostaService(propostaRepo)
	propostaHandler := handler.NewPropostaHandler(propostaSvc)

	authGroup := e.Group("")
	authGroup.Use(middlewares.Autenticacao(usuarioSvc))

	authGroup.GET("/orcamentos", orcamentoHandler.Index)
	authGroup.GET("/orcamentos/novo", orcamentoHandler.Create)
	authGroup.GET("/orcamentos/visualizar/:id", orcamentoHandler.Show)
	authGroup.GET("/orcamentos/editar/:id", orcamentoHandler.Edit)
	authGroup.POST("/orcamentos/salvar", orcamentoHandler.Save)

	authGroup.GET("/propostas/item/:id", propostaHandler.GetPropostasByOrcamentoItemID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Error on starting server")
	}
}
