package updateusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func FindUserByID(db *gorm.DB, userID uint64) (models.Users, error) {
	var user models.Users
	result := db.Where("id = ?", userID).First(&user)
	return user, result.Error
}

func UpdateUser(db *gorm.DB, userID uint64, updateUser models.Users) (models.Users, error) {
	var user models.Users
	result := db.Where("id = ?", userID).Updates(updateUser).First(&user)
	return user, result.Error
}
