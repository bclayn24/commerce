package handlers

import (
	"commerce/db"
	"commerce/templates"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func WatchlistHandler(c echo.Context) error {
	ctx := c.Request().Context()
	listings, err := db.Q.GetUserWatchlist(ctx, ctx.Value("user").(db.User).ID)
	if err != nil {
		return err
	}

	return templates.Render(c, templates.Watchlist(listings))
}

func AddToWatchlistHandler(c echo.Context) error {
	ctx := c.Request().Context()
	listingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user, ok := ctx.Value("user").(db.User)
	if !ok {
		return echo.ErrUnauthorized
	}

	err = db.Q.AddToWatchlist(ctx, db.AddToWatchlistParams{
		UserID:    user.ID,
		ListingID: int64(listingID),
	})

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/watchlist")
}

func RemoveFromWatchlistHandler(c echo.Context) error {
	ctx := c.Request().Context()
	listingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user, ok := ctx.Value("user").(db.User)
	if !ok {
		return echo.ErrUnauthorized
	}

	if err = db.Q.RemoveFromWatchlist(ctx, db.RemoveFromWatchlistParams{
		UserID:    user.ID,
		ListingID: int64(listingID),
	}); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/watchlist")
}
