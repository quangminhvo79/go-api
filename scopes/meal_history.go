package scopes

import (
  "gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/authentication"
)

func MealHistories(db *gorm.DB) *gorm.DB {
  return db.Model(&models.MealHistory{}).Joins("User").Joins("Dish").Joins("Session").
  									 Where("User.email = ?", authentication.Claims.Email)
}

