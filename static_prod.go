//go:build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed static
var assets embed.FS

func static() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.FS(assets)))
}
