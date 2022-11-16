package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name  string
	Price int
	Stock int
}
