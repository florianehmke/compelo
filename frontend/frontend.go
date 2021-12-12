//go:generate go run frontend_generate_models.go
package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed compelo/dist
var embeddedFiles embed.FS

func FileSystem() http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, "compelo/dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
