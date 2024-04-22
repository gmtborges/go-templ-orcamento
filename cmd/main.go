package main

import (
	"github.com/labstack/echo/v4"

	"orcamento-auto/handlers"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.HandleGetIndex)
	e.GET("/politica-privacidade", handlers.HandleGetPolicy)

	e.Logger.Fatal(e.Start(":3000"))
}
