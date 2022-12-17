package database

import (
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
)

func MealHistoriesByUser(uID int) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Model(&models.MealHistory{}).Where("user_id = ?", uID)
  }
}

func MealHistoriesBySession(sID int) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Model(&models.MealHistory{}).Where("session_id = ?", sID)
  }
}

func MealHistoryByID(ID int) *gorm.DB {
  return DB.Model(&models.MealHistory{}).Where("id = ?", ID)
}

func MealHistories() *gorm.DB {
  return DB.Model(&models.MealHistory{}).Joins("User").Joins("Dish").Joins("Session")
}
