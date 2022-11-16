package deleteusers

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gorm.io/gorm"
)

func FindUserByID(db *gorm.DB, userID uint64) (models.Users, error) {
	var user models.Users
	result := db.Where("id = ?", userID).First(&user)
	return user, result.Error
}

func DeleteUser(db *gorm.DB, userID uint64) error {
	var user models.Users
	result := db.Where("id = ?", userID).Delete(&user)
	return result.Error
}
