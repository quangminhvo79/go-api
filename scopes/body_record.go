package scopes

import (
  "gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/authentication"
)

func BodyRecord(db *gorm.DB) *gorm.DB {
	return db.Model(&models.BodyRecord{}).Joins("User").Where("User.email = ?", authentication.Claims.Email)
}
