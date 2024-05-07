package main

import (
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/handlers"
)

func main() {
	e := echo.New()
	e.Any("/*", echo.WrapHandler(public()))
	e.GET("/", handlers.GetIndex)
	e.GET("/politica-privacidade", handlers.GetPolicy)
	e.GET("/auth/login", handlers.GetLogin)

	e.Logger.Fatal(e.Start(":3000"))
}
