package cancelorder

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, orderID uint64, userID uint64) *others.BaseResponse {
	var errResponse *others.BaseResponse

	_, err := FindOrderByID(db, orderID, userID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Order Cannot Found",
			StatusCode: http.StatusNotFound,
		}
		return errResponse
	}

	err = DeleteOrder(db, orderID, userID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Order Cannot Delete",
			StatusCode: http.StatusNotAcceptable,
		}
		return errResponse
	}

	return errResponse
}
