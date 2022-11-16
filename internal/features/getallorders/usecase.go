package getallorders

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, userID uint64) ([]Response, *others.BaseResponse) {
	var errorResponse *others.BaseResponse

	_, err := FindUserByID(db, userID)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "User Cannot Found",
			StatusCode: http.StatusNotFound,
		}
	}

	orders, err := GetAllOrderByUser(db, userID)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Not Found Data",
			StatusCode: http.StatusNotFound,
		}
		return []Response{}, errorResponse
	}

	return orders, errorResponse

}
