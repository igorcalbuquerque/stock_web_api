package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	modelController := controllers.ModelController{DB: db}
	colorController := controllers.ColorController{DB: db}
	productController := controllers.ProductController{DB: db}

	// CRUD Models
	r.GET("/models", modelController.GetModels)
	r.POST("/models", modelController.CreateModel)
	r.PUT("/models/:id", modelController.EditModel)
	r.DELETE("/models/:id", modelController.DeleteModel)
	r.GET("/models/:id", modelController.GetModelById)
	// CRUD Colors
	r.GET("/colors", colorController.GetColors)
	r.POST("/colors", colorController.CreateColor)
	r.PUT("/colors/:id", colorController.EditColor)
	r.DELETE("/colors/:id", colorController.DeleteColor)
	r.GET("/colors/:id", colorController.GetColorById)
	// CRUD Products
	r.GET("/products", productController.GetProducts)
	r.POST("/products", productController.CreateProduct)

	return r
}
