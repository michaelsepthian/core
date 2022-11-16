package createpayment

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, orderID uint, userID uint) (models.Products, models.Payments, *others.BaseResponse) {
	var errorResponse *others.BaseResponse

	order, err := FindOrderByID(db, orderID, userID)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Order Not Found",
			StatusCode: http.StatusNotFound,
		}
		return models.Products{}, models.Payments{}, errorResponse
	}

	product, err := UpdateStockProduct(db, order.ProductID, order.Quantity)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Stock Not Enough",
			StatusCode: http.StatusNotAcceptable,
		}
		return models.Products{}, models.Payments{}, errorResponse
	}

	total := CalculateTotalPayment(order, product)

	newPayment := models.Payments{
		Order: order,
		Total: total,
	}

	payment, err := CreatePayment(db, newPayment)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Failed To Create Payment",
			StatusCode: http.StatusNotAcceptable,
		}
		return models.Products{}, models.Payments{}, errorResponse
	}

	newTransaction := models.Transactions{
		Order:   order,
		Payment: payment,
	}

	_, err = CreateTransaction(db, newTransaction)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Failed To Create Transaction",
			StatusCode: http.StatusNotAcceptable,
		}
		return models.Products{}, models.Payments{}, errorResponse
	}

	return product, payment, errorResponse
}

func CalculateTotalPayment(order models.Orders, product models.Products) int {
	return order.Quantity * product.Price
}
