package web

import (
	"fmt"
	"net/http"

	"github.com/potato-squad-org/site/web/handler"
)

func HTTPServer(port int) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler.Handler(),
	}
}
