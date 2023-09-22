package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ModelController struct {
	DB *gorm.DB
}

func (modelController *ModelController) GetModels(c *gin.Context) {
	var models []models.Model
	modelController.DB.Find(&models)
	c.JSON(http.StatusOK, models)
}

func (modelController *ModelController) CreateModel(c *gin.Context) {
	var newModel models.Model

	if err := c.ShouldBindJSON(&newModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := modelController.DB.Create(&newModel)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create a new model",
		})
		return
	}

	c.JSON(http.StatusCreated, newModel)
}

func (modelController *ModelController) EditModel(c *gin.Context) {
	modelID := c.Param("id")

	var existingModel models.Model
	result := modelController.DB.First(&existingModel, "id = ?", modelID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Model not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&existingModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result = modelController.DB.Save(&existingModel)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update the model",
		})
		return
	}
	c.JSON(http.StatusOK, existingModel)
}

func (modelController *ModelController) DeleteModel(c *gin.Context) {
	modelID := c.Param("id")

	var existingModel models.Model
	result := modelController.DB.First(&existingModel, "id = ?", modelID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Model not found",
		})
		return
	}

	result = modelController.DB.Delete(&existingModel)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete the model",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Model deleted successfully",
	})
}

func (modelController *ModelController) GetModelById(c *gin.Context) {
	modelID := c.Param("id")

	var existingModel models.Model
	result := modelController.DB.First(&existingModel, "id = ?", modelID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Model not found",
		})
		return
	}

	c.JSON(http.StatusOK, existingModel)
}
