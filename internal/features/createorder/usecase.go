package createorder

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, userID uint, productID uint, quantity int) (models.Orders, *others.BaseResponse) {
	var errorResponse *others.BaseResponse

	user, err := FindUserByID(db, userID)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Failed To Find User",
			StatusCode: http.StatusNotFound,
		}
		return models.Orders{}, errorResponse
	}

	product, err := FindProductByID(db, productID, quantity)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Failed To Get Product",
			StatusCode: http.StatusNotFound,
		}
		return models.Orders{}, errorResponse
	}

	newOrder := models.Orders{
		User:     user,
		Product:  product,
		Quantity: quantity,
	}

	order, err := CreateOrder(db, newOrder)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Failed To Save Order",
			StatusCode: http.StatusNotFound,
		}
		return models.Orders{}, errorResponse
	}

	return order, errorResponse
}
