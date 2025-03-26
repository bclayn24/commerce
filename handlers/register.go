package handlers

import (
	"commerce/db"
	"commerce/templates"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHandler(c echo.Context) error {
	if c.Request().Method == "POST" {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username == "" || password == "" {
			return c.String(http.StatusBadRequest, "username or password is empty")
		}

		ctx := c.Request().Context()
		if user, err := db.Q.GetUser(ctx, username); err == nil && user.Username == username {
			return c.String(http.StatusBadRequest, "username already exists")
		}

		// hash password
		hash := sha256.Sum256([]byte(password))
		hashStr := fmt.Sprintf("%x", hash)

		userId, err := db.Q.CreateUser(ctx, db.CreateUserParams{
			Username: username,
			Password: hashStr,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to create user")
		}

		err = createSession(userId, c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to create session")
		}

		return c.Redirect(http.StatusFound, "/")

	}

	return templates.Render(c, templates.Register())
}
