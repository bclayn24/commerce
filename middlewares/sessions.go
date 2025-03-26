package middlewares

import (
	"commerce/db"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Session(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Path()
		if path == "/login" || path == "/register" {
			return next(c)
		}

		tokenCookie, err := c.Cookie("session_token")

		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				return next(c)
			}
			// For any other type of error, return a bad request status
			return next(c)
		}
		token := tokenCookie.Value

		ctx := c.Request().Context()
		userSession, err := db.Q.GetSession(ctx, token)

		if err != nil {
			// If the session token is not present in session map, return an unauthorized error
			return next(c)
		}

		// If the session is present, but has expired, we can delete the session, and return
		// an unauthorized status
		if userSession.IsExpired() {
			db.Q.DeleteSession(ctx, token)
			return next(c)
		}

		user, err := db.Q.GetUserById(ctx, userSession.UserID)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to get user")
		}

		ctx = context.WithValue(c.Request().Context(), "user", user)
		nr := c.Request().WithContext(ctx)
		c.SetRequest(nr)

		return next(c)
	}
}
