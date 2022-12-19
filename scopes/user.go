package scopes

import (
  "gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/authentication"
)

func User(db *gorm.DB) *gorm.DB {
	return db.Model(&models.User{}).Where("email = ?", authentication.Claims.Email)
}
