package handlers

import (
	"commerce/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CommentHandler(c echo.Context) error {
	ctx := c.Request().Context()

	comment := c.FormValue("comment_text")

	listing_id, err := strconv.Atoi(c.FormValue("listing_id"))
	if err != nil {
		return err
	}

	user, ok := ctx.Value("user").(db.User)
	if !ok {
		return c.String(http.StatusInternalServerError, "failed to get user")
	}

	db.Q.CreateComment(ctx, db.CreateCommentParams{
		ListingID: int64(listing_id),
		OwnerID:   user.ID,
		OwnerName: user.Username,
		Comment:   comment,
	})

	return c.Redirect(302, "/listing/"+c.FormValue("listing_id"))
}
