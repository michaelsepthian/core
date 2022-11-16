package createproducts

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, products models.Products) (models.Products, error) {
	result := db.Create(&products)
	return products, result.Error
}
