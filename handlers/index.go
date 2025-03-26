package handlers

import (
	"commerce/db"
	"commerce/templates"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {

	listings, err := db.Q.GetListings(c.Request().Context())
	if err != nil {
		return c.String(500, "failed to get listings")
	}

	return templates.Render(c, templates.Index(listings))
}
