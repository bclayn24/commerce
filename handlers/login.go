package handlers

import (
	"commerce/db"
	"commerce/templates"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	if c.Request().Method == "POST" {
		username := c.FormValue("username")
		password := c.FormValue("password")
		if username == "" || password == "" {
			return c.String(400, "username or password is empty")
		}

		ctx := c.Request().Context()

		// Get the expected password from our in memory map
		user, err := db.Q.GetUser(ctx, username)

		hashStr := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

		if err != nil || user.Password != hashStr {
			return c.String(http.StatusUnauthorized, "wrong username or password")
		}

		err = createSession(user.ID, c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to create session")
		}

		return c.Redirect(http.StatusFound, "/")

	}

	return templates.Render(c, templates.Login())
}

func LogoutHandler(c echo.Context) error {
	sessionCookie, _ := c.Cookie("session_token")
	db.Q.DeleteSession(c.Request().Context(), sessionCookie.Value)
	return c.Redirect(http.StatusFound, "/")
}
