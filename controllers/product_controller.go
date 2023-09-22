package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func (productController *ProductController) GetProducts(c *gin.Context) {
	var products []models.Product
	// productController.DB.Preload("Model").Preload("Color").Find(&products)
	productController.DB.Find(&products)
	// productController.DB.Select("id, imei, active, model_id, color_id").Find(&products)
	c.JSON(http.StatusOK, products)
}

func (productController *ProductController) CreateProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var existingColor models.Color
	resultColor := productController.DB.First(&existingColor, "id = ?", newProduct.ColorID)

	if resultColor.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Color not found. The color is necessary for create a product",
		})
		return
	}

	var existingModel models.Model
	resultModel := productController.DB.First(&existingModel, "id = ?", newProduct.ModelID)

	if resultModel.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Model not found. The model is necessary for create a product",
		})
		return
	}

	result := productController.DB.Create(&newProduct)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create new color",
		})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}
