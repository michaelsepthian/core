package finishtransaction

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func CheckTransaction(db *gorm.DB, transactionID uint64) (models.Transactions, error) {
	var transaction models.Transactions
	result := db.Where("id = ? AND status_payment = 'UNPAID'", transactionID).First(&transaction)
	return transaction, result.Error
}

func FindTransactionByID(db *gorm.DB, transactionID uint64) (models.Transactions, error) {
	var transaction models.Transactions
	result := db.Where("id = ?", transactionID).First(&transaction)
	return transaction, result.Error
}

func UpdateTransaction(db *gorm.DB, transactionID uint64, method string) (models.Transactions, error) {
	var transaction models.Transactions
	result := db.Model(&transaction).Where("id = ?", transactionID).Updates(map[string]interface{}{"status_payment": "PAID", "method_payment": method}).First(&transaction)
	return transaction, result.Error
}

func UpdateOrder(db *gorm.DB, transactionID uint64) error {
	var order models.Orders

	findTransaction, _ := FindTransactionByID(db, transactionID)
	result := db.Model(&order).Where("id = ?", findTransaction.OrderID).Update("status", "PAID").First(&order)
	return result.Error
}

func GetTransaction(db *gorm.DB, transactionID uint64) (Response, error) {
	var response Response
	result := db.Table("transactions").
		Select("transactions.status_payment as status_payment, transactions.method_payment as method, prd.name as name_product, ord.quantity as quantity, py.total as total").
		Joins("JOIN payments as py on py.id = transactions.payment_id").
		Joins("JOIN orders as ord on ord.id = transactions.order_id").
		Joins("JOIN products as prd on prd.id = ord.product_id").
		Where("transactions.id = ?", transactionID).
		Find(&response)
	return response, result.Error
}
