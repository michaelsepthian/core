package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	UserID    uint `json:"user_id" gorm:"index"`
	ProductID uint `json:"product_id" gorm:"index"`
	Quantity  int
	Status    string   `gorm:"default:UNPAID"`
	User      Users    `gorm:"foreignKey:UserID"`
	Product   Products `gorm:"foreignKey:ProductID"`
}
