//go:build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var assets embed.FS

func public() http.Handler {
	return http.FileServer(http.FS(assets))
}

