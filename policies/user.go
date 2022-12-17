package policies

import (
  "gorm.io/gorm"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/models"
)

func UserScope(email string) *gorm.DB {
	return database.DB.Model(&models.User{}).Where("email = ?", email)
}
