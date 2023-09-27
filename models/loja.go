package models

import "gorm.io/gorm"

type Loja struct {
	gorm.Model
	Name   string
	Active bool
	Cidade string
}
