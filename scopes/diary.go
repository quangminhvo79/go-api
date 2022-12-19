package scopes

import (
  "gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/authentication"
)

func Diary(db *gorm.DB) *gorm.DB {
	return db.Model(&models.Diary{}).Joins("User").Where("User.email = ?", authentication.Claims.Email)
}
