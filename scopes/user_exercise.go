package scopes

import (
  "gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/authentication"
)

func UserExercise(db *gorm.DB) *gorm.DB {
	return db.Model(&models.UserExercise{}).Joins("User").Where("User.email = ?", authentication.Claims.Email)
}
