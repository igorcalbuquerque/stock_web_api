package models

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name   string
	Active bool
}
