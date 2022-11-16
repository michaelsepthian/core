package createpayment

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func FindOrderByID(db *gorm.DB, orderID uint, userID uint) (models.Orders, error) {
	var order models.Orders
	result := db.Where("id = ? AND user_id = ?", orderID, userID).First(&order)
	return order, result.Error
}

func CreatePayment(db *gorm.DB, payments models.Payments) (models.Payments, error) {
	result := db.Create(&payments)
	return payments, result.Error
}

func UpdateStockProduct(db *gorm.DB, productID uint, quantity int) (models.Products, error) {
	var product models.Products
	result := db.Model(&product).Where("id = ?", productID).Update("stock", gorm.Expr("stock - ?", quantity)).Find(&product)
	return product, result.Error
}

func CreateTransaction(db *gorm.DB, transaction models.Transactions) (models.Transactions, error) {
	result := db.Create(&transaction)
	return transaction, result.Error
}
