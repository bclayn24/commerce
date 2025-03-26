package handlers

import (
	"commerce/db"
	"commerce/templates"
	"database/sql"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CategoriesHandler(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := db.Q.GetCategories(ctx)

	if err != nil {
		return err
	}

	return templates.Render(c, templates.Categories(categories))
}

func CategoryListingsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	categoryId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	listings, err := db.Q.GetListingsByCategoryId(ctx, sql.NullInt64{
		Valid: true,
		Int64: int64(categoryId),
	})

	if err != nil {
		return err
	}

	category, err := db.Q.GetCategoryById(ctx, int64(categoryId))

	if err != nil {
		return err
	}

	return templates.Render(c, templates.CategoryListings(category, listings))
}
