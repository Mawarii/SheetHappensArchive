package controller

import (
	"fmt"
	"net/http"
	"sheethappens/backend/database"
	"sheethappens/backend/model"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	return c.File("frontend/html/register.html")
}

func RegisterUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var existingUser model.User
	result := database.DB().Where("username = ?", username).First(&existingUser)
	if result.Error == nil {
		return c.String(http.StatusConflict, "Username already taken")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Password hashing failed")
	}

	newUser := model.User{
		Username: username,
		Password: hashedPassword,
	}

	if err := database.DB().Create(&newUser).Error; err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return c.String(http.StatusInternalServerError, "Registration failed")
	}

	return c.String(http.StatusCreated, "Registration successful")
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func Login(c echo.Context) error {
	return c.File("frontend/html/login.html")
}

func Authenticate(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User
	result := database.DB().Where("username = ?", username).First(&user)

	if result.Error == nil && comparePassword(user.Password, password) {
		// Authentifizierung erfolgreich

		// Erstellen Sie oder aktualisieren Sie die Sitzung mit dem Benutzernamen
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		// Speichern Sie den Benutzernamen in der Sitzung
		sess.Values["username"] = username

		if err := sess.Save(c.Request(), c.Response()); err != nil {
			// Behandeln Sie den Fehler, wenn die Sitzung nicht gespeichert werden kann
			return err
		}

		// Nach erfolgreicher Authentifizierung weiterleiten
		return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("character"))
	}

	// Authentifizierung fehlgeschlagen
	return c.String(http.StatusUnauthorized, "Invalid login")
}
