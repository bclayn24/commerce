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

func CreateListingHandler(c echo.Context) error {
	ctx := c.Request().Context()

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

		user, ok := ctx.Value("user").(db.User)
		if !ok {
			return c.String(http.StatusInternalServerError, "failed to get user")
		}

		db.Q.CreateListing(ctx, db.CreateListingParams{
			Title:         title,
			Description:   sql.NullString{String: description, Valid: true},
			StartingPrice: startingPrice,
			ImageUrl:      image,
			OwnerID:       user.ID,
			CreatedAt:     sql.NullTime{Time: time.Now().UTC(), Valid: true},
			CategoryID:    catId,
		})

		return c.Redirect(http.StatusFound, "/")
	}

	categories, err := db.Q.GetCategories(ctx)
	if err != nil {
		return err
	}

	return templates.Render(c, templates.CreateListing(categories))
}
