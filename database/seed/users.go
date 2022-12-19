package seed

import (
	"errors"
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
)

func CreateUsers() {
  if DB.Migrator().HasTable(&models.User{}) {
    if err := DB.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      var users = []models.User{
        {Email: "man@gmail.com", Password: "12345", Username: "Man", AchievementWeightFrom: 60, AchievementWeightTo: 75},
        {Email: "woman@gmail.com", Password: "12345", Username: "Woman", AchievementWeightFrom: 60, AchievementWeightTo: 49},
      }

      DB.Create(&users)
    }
  }
}
