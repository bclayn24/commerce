package handlers

import (
	"commerce/db"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BidHandler(c echo.Context) error {

	ctx := c.Request().Context()

	listing_id, err := strconv.Atoi(c.FormValue("listing_id"))
	if err != nil {
		return err
	}

	max_bid, err := strconv.ParseFloat(c.FormValue("max_bid"), 64)
	if err != nil {
		return err
	}

	listing, err := db.Q.GetListing(ctx, int64(listing_id))
	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		return err
	}

	user, ok := ctx.Value("user").(db.User)
	if !ok {
		return errors.New("failed to get user")
	}

	if amount > listing.StartingPrice && amount > max_bid {
		err = db.Q.CreateBid(ctx, db.CreateBidParams{
			Amount:    amount,
			ListingID: int64(listing_id),
			UserID:    user.ID,
			OwnerName: user.Username,
		})
		if err != nil {
			return err
		}
	}

	return c.Redirect(302, "/listing/"+c.FormValue("listing_id"))
}
