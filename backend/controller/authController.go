package controller

import (
	"fmt"
	"net/http"
	"sheethappens/backend/database"
	"sheethappens/backend/model"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c echo.Context) error {
	return c.File("frontend/html/register.html")
}

func RegisterUser(c echo.Context) error {
	// Read user data from the request
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Check if the user already exists
	var existingUser model.User
	result := database.DB().Where("username = ?", username).First(&existingUser)
	if result.Error == nil {
		return c.String(http.StatusConflict, "Username already taken")
	}

	// Hash the password using Argon2
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Password hashing failed")
	}

	// Create a new user and save it in the database
	newUser := model.User{
		Username: username,
		Password: hashedPassword,
	}

	if err := database.DB().Create(&newUser).Error; err != nil {
		// Log the error for debugging
		fmt.Printf("Error creating user: %v\n", err)
		return c.String(http.StatusInternalServerError, "Registration failed")
	}

	return c.String(http.StatusCreated, "Registration successful")
}

func hashPassword(password string) (string, error) {
	// Generieren Sie einen Hash des Passworts mit bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePassword(hashedPassword string, password string) bool {
	// Vergleichen Sie das gehashte Passwort mit dem eingegebenen Passwort
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
	if result.Error == gorm.ErrRecordNotFound {
		return c.String(http.StatusUnauthorized, "Invalid login")
	} else if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Database error")
	}

	if !comparePassword(user.Password, password) {
		return c.String(http.StatusUnauthorized, "Invalid login")
	}

	return c.String(http.StatusOK, "Login successful")
}
