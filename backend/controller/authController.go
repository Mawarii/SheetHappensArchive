package controller

import (
	"net/http"
	"sheethappens/backend/database"
	"sheethappens/backend/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	return c.File("frontend/html/login.html")
}

func Authenticate(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User
	result := database.DB().Where("username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return c.String(http.StatusUnauthorized, "Invalid login")
	} else if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Database error")
	}

	// Einfacher Passwort-Check: Überprüfen Sie, ob das eingegebene Passwort dem in der Datenbank gespeicherten Passwort entspricht
	if password != user.Password {
		return c.String(http.StatusUnauthorized, "Invalid login")
	}

	// Hier können Sie eine Sitzung starten oder einen Token generieren, um den Benutzer anzumelden

	return c.String(http.StatusOK, "Login successful")
}
