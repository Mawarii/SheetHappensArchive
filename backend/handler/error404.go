package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFoundHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().WriteHeader(http.StatusNotFound)

		return c.File("frontend/html/error404.html")
	}
}
