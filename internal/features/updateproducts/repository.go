package updateproducts

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func FindProductByID(db *gorm.DB, productID uint64) (models.Products, error) {
	var product models.Products
	result := db.Where("id = ?", productID).First(&product)
	return product, result.Error
}

func UpdateProduct(db *gorm.DB, productID uint64, updateProduct models.Products) (models.Products, error) {
	var product models.Products
	result := db.Where("id = ?", productID).Updates(updateProduct).First(&product)
	return product, result.Error
}
