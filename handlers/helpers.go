package handlers

import (
	"commerce/db"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func createSession(userId int64, c echo.Context) error {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Minute)

	err := db.Q.CreateSession(c.Request().Context(), db.CreateSessionParams{
		UserID:    userId,
		Token:     sessionToken,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		return err
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	return nil
}
