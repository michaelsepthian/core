package getallorders

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

type listOrders struct {
	NameProduct string `json:"nameProduct"`
}

func FindUserByID(db *gorm.DB, userID uint64) (models.Orders, error) {
	var order models.Orders
	result := db.Where("id = ?", userID).First(&order)
	return order, result.Error
}

func GetAllOrderByUser(db *gorm.DB, userID uint64) ([]Response, error) {
	var response []Response
	result := db.Table("orders").Select("orders.quantity as quantity, prd.name as name_product").
		Joins("JOIN products as prd on prd.id = orders.product_id").
		Where("orders.user_id = ? AND orders.status = ?", userID, "UNPAID").
		Scan(&response)
	return response, result.Error
}
