package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var handler = func() *echo.Echo {
	e := echo.New()
	e.Static("/", "public")
	return e
}()

func Handler() http.Handler {
	return handler
}
