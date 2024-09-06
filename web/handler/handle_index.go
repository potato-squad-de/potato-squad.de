package handler

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/potato-squad-org/site/web/ui"
	"github.com/potato-squad-org/site/web/ui/shared"
)

func init() {
	handler.GET("/", func(c echo.Context) error {
		page := shared.Page(shared.PageProps{
			Title: "Home",
			Body:  ui.Index(),
		})
		return c.(*customContext).NodeAsHTML(http.StatusOK, page)
	})
}
