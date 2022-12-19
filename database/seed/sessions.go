package seed

import (
  "errors"
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
)

func CreateSessions() {
  if DB.Migrator().HasTable(&models.Session{}) {
    if err := DB.First(&models.Session{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      var sessions = []models.Session{{Name: "Morning"}, {Name: "Lunch"}, {Name: "Dinner"}, {Name: "Snack"}}
      DB.Create(&sessions)
    }
  }
}
