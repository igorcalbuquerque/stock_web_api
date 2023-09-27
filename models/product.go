package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Imei       string
	Active     bool
	ModelID    uint
	ColorID    uint
	EstoqueQtd int32
	Capacidade int32
	Observacao string
	Bateria    int32
}
