package deleteusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, userID uint64) *others.BaseResponse {
	var errResponse *others.BaseResponse

	_, err := FindUserByID(db, userID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "User Cannot Found",
			StatusCode: http.StatusNotFound,
		}
		return errResponse
	}

	err = DeleteUser(db, userID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "User Cannot Delete",
			StatusCode: http.StatusNotAcceptable,
		}
		return errResponse
	}

	return errResponse
}
