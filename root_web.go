package root

import (
	"embed"
	"net/http"
)

type web struct {
	staticRootDir string
	staticFS      http.FileSystem
}

//go:embed web/static
var embeddedStaticFS embed.FS

var w = &web{
	staticRootDir: "web/static",
	staticFS:      http.FS(embeddedStaticFS),
}

func Web() *web {
	return w
}

func (w *web) StaticRootDir() string {
	return w.staticRootDir
}

func (w *web) StaticFS() http.FileSystem {
	return w.staticFS
}
