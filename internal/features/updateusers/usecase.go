package updateusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, userID uint64, name string, age int, address string) (models.Users, *others.BaseResponse) {
	var errResponse *others.BaseResponse

	_, err := FindUserByID(db, userID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "User Cannot Found",
			StatusCode: http.StatusNotFound,
		}
		return models.Users{}, errResponse
	}

	updateUser := models.Users{
		Name:    name,
		Age:     age,
		Address: address,
	}

	user, err := UpdateUser(db, userID, updateUser)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Failed To Update Data User",
			StatusCode: http.StatusBadRequest,
		}
		return models.Users{}, errResponse
	}

	return user, errResponse
}
