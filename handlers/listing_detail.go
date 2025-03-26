package handlers

import (
	"commerce/db"
	"commerce/models"
	"commerce/templates"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListingHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	listing, err := db.Q.GetListing(ctx, int64(id))
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

	owner, err := db.Q.GetUserById(ctx, listing.OwnerID)
	if err != nil {
		return err
	}

	isInWatchlist := false
	user, ok := ctx.Value("user").(db.User)
	if ok {
		_, err = db.Q.IsInWatchlist(ctx, db.IsInWatchlistParams{
			UserID:    user.ID,
			ListingID: listing.ID,
		})

		isInWatchlist = err == nil
	}

	max_bid, _ := db.Q.GetMaxBid(ctx, listing.ID)

	bids, err := db.Q.GetBidsByListingId(ctx, listing.ID)
	if err != nil {
		return err
	}

	comments, err := db.Q.GetCommentsByListingId(ctx, listing.ID)
	if err != nil {
		return err
	}

	return templates.Render(c, templates.ListingDetail(models.ListingDetailArgs{
		Listing:       listing,
		Owner:         owner,
		Category:      category,
		Bids:          bids,
		MaxBid:        max_bid,
		IsInWatchlist: isInWatchlist,
		Comments:      comments,
	}))
}

func CloseListingHandler(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	err = db.Q.CloseListing(ctx, int64(id))

	if err != nil {
		return err
	}

	return c.Redirect(302, "/")
}
