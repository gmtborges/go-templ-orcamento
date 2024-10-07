//go:build dev

package static

import (
	"fmt"
	"net/http"
	"os"
)

func Serve() http.Handler {
	fmt.Println("building static files for development...")
	return http.StripPrefix("/static/", http.FileServerFS(os.DirFS("static")))
}
