package handler

import (
	"bytes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	g "github.com/maragudk/gomponents"
	root "github.com/potato-squad-org/site"
)

var handler = echo.New()

func init() {
	handler.Use(middleware.CORS())
	handler.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	handler.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       root.Web().StaticRootDir(),
		Filesystem: root.Web().StaticFS(),
	}))
}

type customContext struct {
	echo.Context
}

func (cc *customContext) NodeAsHTML(status int, node g.Node) error {
	var buf bytes.Buffer
	err := node.Render(&buf)
	if err != nil {
		return err
	}
	return cc.HTML(status, buf.String())
}

func init() {
	handler.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customContext{Context: c}
			return next(cc)
		}
	})
}

func Handler() http.Handler {
	return handler
}
