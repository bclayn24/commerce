package handlers

import (
	"commerce/db"
	"commerce/templates"
	"errors"

	"github.com/labstack/echo/v4"
)

func MyListingsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := ctx.Value("user").(db.User)
	if !ok {
		return errors.New("failed to get user")
	}

	listings, err := db.Q.GetListingsByOwnerId(ctx, user.ID)
	if err != nil {
		return err
	}

	return templates.Render(c, templates.MyListings(listings))
}
