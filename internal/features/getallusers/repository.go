package getallusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.Users, error) {
	var users []models.Users
	result := db.Find(&users)
	return users, result.Error
}
