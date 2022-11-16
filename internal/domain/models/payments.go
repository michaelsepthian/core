package models

import (
	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model
	OrderID uint `json:"order_id" gorm:"index"`
	Total   int
	Order   Orders `gorm:"foreignKey:OrderID"`
}
