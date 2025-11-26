package handlers

import (
	"net/http"
	"strconv"

	"github.com/minseoi/gorizon/db"
	"github.com/minseoi/gorizon/models"

	"github.com/labstack/echo/v4"
)

// GetIngredients - 모든 식재료 조회
func GetIngredients(c echo.Context) error {
	var ingredients []models.Ingredient

	result := db.DB.Find(&ingredients)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch ingredients",
		})
	}

	return c.JSON(http.StatusOK, ingredients)
}

// CreateIngredient - 식재료 추가
func CreateIngredient(c echo.Context) error {
	var newIngredient models.Ingredient

	if err := c.Bind(&newIngredient); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Validation
	if newIngredient.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name is required",
		})
	}

	result := db.DB.Create(&newIngredient)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create ingredient",
		})
	}

	return c.JSON(http.StatusCreated, newIngredient)
}

// UpdateIngredient - 식재료 이름 변경
func UpdateIngredient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ingredient ID",
		})
	}

	var ingredient models.Ingredient
	result := db.DB.First(&ingredient, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Ingredient not found",
		})
	}

	var updateData models.Ingredient
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Validation
	if updateData.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name is required",
		})
	}

	// Update only the name field
	ingredient.Name = updateData.Name

	result = db.DB.Save(&ingredient)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update ingredient",
		})
	}

	return c.JSON(http.StatusOK, ingredient)
}

// DeleteIngredient - 식재료 제거
func DeleteIngredient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ingredient ID",
		})
	}

	var ingredient models.Ingredient
	result := db.DB.First(&ingredient, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Ingredient not found",
		})
	}

	result = db.DB.Delete(&ingredient)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete ingredient",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Ingredient deleted successfully",
	})
}
