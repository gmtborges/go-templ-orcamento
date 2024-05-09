//go:build dev

package main

import (
	"fmt"
	"net/http"
	"os"
)

func static() http.Handler {
	fmt.Println("building static files for development...")
	return http.StripPrefix("/static/", http.FileServerFS(os.DirFS("static")))
}
