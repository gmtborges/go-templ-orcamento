//go:build !dev

package static

import (
	"embed"
	"net/http"
)

//go:embed static
var static embed.FS

func Serve() http.Handler {
	return http.FileServer(http.FS(static))
}
