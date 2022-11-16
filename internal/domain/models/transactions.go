package models

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	OrderID       uint   `json:"order_id" gorm:"index"`
	PaymentID     uint   `json:"payment_id" gorm:"index"`
	StatusPayment string `gorm:"default:UNPAID"`
	MethodPayment string
	Order         Orders   `gorm:"foreignKey:OrderID"`
	Payment       Payments `gorm:"foreignKey:PaymentID"`
}
