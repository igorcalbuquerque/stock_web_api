package database

import (
	"api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "user=admin dbname=teste password=123456 host=localhost port=5432 sslmode=disable"

	// Open a database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Model{})
	db.AutoMigrate(&models.Color{})
	db.AutoMigrate(&models.Product{})

	return db, nil
}
