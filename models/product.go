package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Imei    string
	Active  bool
	ModelID uint
	ColorID uint
}
