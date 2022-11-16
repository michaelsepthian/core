package createusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, users models.Users) (models.Users, error) {
	result := db.Create(&users)
	return users, result.Error
}
