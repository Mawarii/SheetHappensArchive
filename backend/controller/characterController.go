package controller

import (
	"net/http"
	"sheethappens/backend/database"
	"sheethappens/backend/model"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreateCharacter(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userID, _ := sess.Values["userID"].(int)

	b := new(model.Character)
	db := database.DB()

	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"error": "Invalid data provided.",
		}

		return c.JSON(http.StatusBadRequest, data)
	}

	character := &model.Character{
		UserID: userID,
		Name:   b.Name,
		Race:   b.Race,
	}

	if err := db.Create(&character).Error; err != nil {
		data := map[string]interface{}{
			"error": "Failed to create a character.",
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": character,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateCharacter(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userID, _ := sess.Values["userID"].(int)

	id := c.Param("id")
	b := new(model.Character)
	db := database.DB()

	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"error": "Invalid data provided.",
		}

		return c.JSON(http.StatusBadRequest, data)
	}

	existingChar := new(model.Character)

	if err := db.Where("user_id = ?", userID).First(&existingChar, id).Error; err != nil {
		data := map[string]interface{}{
			"error": "Character not found.",
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existingChar.Name = b.Name
	existingChar.Race = b.Race
	if err := db.Save(&existingChar).Error; err != nil {
		data := map[string]interface{}{
			"error": "Failed to update the character.",
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existingChar,
	}

	return c.JSON(http.StatusOK, response)
}

func GetCharacter(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userID, _ := sess.Values["userID"].(int)

	id := c.Param("id")
	db := database.DB()

	var chars []*model.Character

	if res := db.Where("user_id = ?", userID).Find(&chars, id); res.Error != nil {
		data := map[string]interface{}{
			"error": "Failed to retrieve characters.",
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	if len(chars) == 0 {
		return c.JSON(http.StatusNotFound, "No characters found.")
	}

	response := map[string]interface{}{
		"data": chars[0],
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteCharacter(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userID, _ := sess.Values["userID"].(int)

	id := c.Param("id")
	db := database.DB()

	character := new(model.Character)

	err := db.Where("user_id = ?", userID).Delete(&character, id).Error
	if err != nil {
		data := map[string]interface{}{
			"error": "Failed to delete the character.",
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Character has been deleted successfully.",
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllCharacters(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userID, _ := sess.Values["userID"].(int)

	db := database.DB()

	var chars []*model.Character

	if res := db.Where("user_id = ?", userID).Find(&chars); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	if len(chars) == 0 {
		return c.JSON(http.StatusOK, "Keine Chars")
	}

	response := map[string]interface{}{
		"data": chars,
	}

	return c.JSON(http.StatusOK, response)
}
