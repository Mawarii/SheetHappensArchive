package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CheckSessionToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		userID, ok := sess.Values["userID"].(uint)

		if ok && userID != 0 {
			return next(c)
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}
}
