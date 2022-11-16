package createorder

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB, orders models.Orders) (models.Orders, error) {
	result := db.Create(&orders)
	return orders, result.Error
}

func FindUserByID(db *gorm.DB, userID uint) (models.Users, error) {
	var user models.Users
	result := db.Where("id = ?", userID).First(&user)
	return user, result.Error
}

func FindProductByID(db *gorm.DB, productID uint, quantity int) (models.Products, error) {
	var product models.Products
	result := db.Where("id = ? AND stock > ?", productID, quantity).First(&product)
	return product, result.Error
}
