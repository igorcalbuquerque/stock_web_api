package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ColorController struct {
	DB *gorm.DB
}

func (colorController *ColorController) GetColors(c *gin.Context) {
	var colors []models.Color
	colorController.DB.Find(&colors)
	c.JSON(http.StatusOK, colors)
}

func (colorController *ColorController) CreateColor(c *gin.Context) {
	var newColor models.Color

	if err := c.ShouldBindJSON(&newColor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	result := colorController.DB.Create(&newColor)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create a new color",
		})
		return
	}

	c.JSON(http.StatusCreated, newColor)
}

func (colorController *ColorController) EditColor(c *gin.Context) {
	modelId := c.Param("id")

	var existingColor models.Color
	result := colorController.DB.First(&existingColor, "id = ?", modelId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Color not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&existingColor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result = colorController.DB.Save(&existingColor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update the color",
		})
		return
	}
	c.JSON(http.StatusOK, existingColor)
}

func (colorController *ColorController) DeleteColor(c *gin.Context) {
	modelID := c.Param("id")

	var existingColor models.Color
	result := colorController.DB.First(&existingColor, "id = ?", modelID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Color not found",
		})
		return
	}

	result = colorController.DB.Delete(&existingColor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete color",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Color deleted successfuly",
	})
}

func (colorController *ColorController) GetColorById(c *gin.Context) {
	modelID := c.Param("id")

	var existingColor models.Color
	result := colorController.DB.First(&existingColor, "id = ?", modelID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Color not found",
		})
		return
	}
	c.JSON(http.StatusOK, existingColor)
}
