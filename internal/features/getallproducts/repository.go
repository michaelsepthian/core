package getallproducts

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB) ([]models.Products, error) {
	var products []models.Products
	result := db.Find(&products)
	return products, result.Error
}
