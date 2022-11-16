package cancelorder

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func FindOrderByID(db *gorm.DB, orderID uint64, userID uint64) (models.Orders, error) {
	var order models.Orders
	result := db.Where("id = ? AND user_id = ?", orderID, userID).First(&order)
	return order, result.Error
}

func DeleteOrder(db *gorm.DB, orderID uint64, userID uint64) error {
	var order models.Orders
	result := db.Where("id = ? AND user_id = ?", orderID, userID).Delete(&order)
	return result.Error
}
