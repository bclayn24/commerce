package handlers

import (
	"commerce/db"
	"commerce/templates"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func EditListingHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	listing, err := db.Q.GetListing(ctx, int64(id))
	if err != nil {
		return err
	}
	owner, err := db.Q.GetUserById(ctx, listing.OwnerID)
	if err != nil {
		return err
	}
	bids, err := db.Q.GetBidsByListingId(ctx, listing.ID)
	if err != nil {
		return err
	}

	var category db.Category
	if listing.CategoryID.Valid {
		category, err = db.Q.GetCategoryById(ctx, listing.CategoryID.Int64)
		if err != nil {
			return err
		}
	}
	categories, err := db.Q.GetCategories(ctx)
	if err != nil {
		return err
	}

	if c.Request().Method == "POST" {
		title := c.FormValue("title")
		description := c.FormValue("description")
		price := c.FormValue("price")
		image := c.FormValue("image")
		categoryId, err := strconv.Atoi(c.FormValue("category"))
		if err != nil {
			return err
		}
		catId := sql.NullInt64{Int64: int64(categoryId), Valid: true}

		if categoryId == 0 {
			catId = sql.NullInt64{Valid: false}
		}

		if title == "" || price == "" || image == "" {
			return c.String(http.StatusBadRequest, "title, image or price is empty")
		}

		startingPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "price is not a valid number")
		}

		db.Q.EditListing(ctx, db.EditListingParams{
			Title:         title,
			Description:   sql.NullString{String: description, Valid: true},
			StartingPrice: startingPrice,
			ImageUrl:      image,
			CategoryID:    catId,
			CreatedAt:     sql.NullTime{Time: time.Now().UTC(), Valid: true},
			ID:            int64(id),
		})

		return c.Redirect(http.StatusFound, "/")
	}
	isInWatchlist := false
	return templates.Render(c, templates.EditListing(categories, listing, owner, category, bids, isInWatchlist))
}
