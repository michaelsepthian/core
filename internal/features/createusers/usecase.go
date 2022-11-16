package createusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, name string, age int, address string) (models.Users, *others.BaseResponse) {
	var errorResponse *others.BaseResponse
	newUsers := models.Users{
		Name:    name,
		Age:     age,
		Address: address,
	}

	user, err := CreateUser(db, newUsers)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Failed To Save User",
			StatusCode: http.StatusInternalServerError,
		}
		return models.Users{}, errorResponse
	}

	return user, errorResponse
}
