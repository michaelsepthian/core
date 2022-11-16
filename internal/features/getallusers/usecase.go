package getallusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB) ([]Response, *others.BaseResponse) {
	var errorResponse *others.BaseResponse

	users, err := GetAllUsers(db)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Empty Data",
			StatusCode: http.StatusNotFound,
		}
		return []Response{}, errorResponse
	}

	responseList := MapUsersToResponse(users)

	return responseList, errorResponse

}

func MapUsersToResponse(users []models.Users) []Response {
	var responselist []Response
	for _, user := range users {
		allUsers := Response{
			Name:    user.Name,
			Age:     user.Age,
			Address: user.Address,
		}
		responselist = append(responselist, allUsers)
	}
	return responselist
}
