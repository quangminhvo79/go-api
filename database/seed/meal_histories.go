package seed

import (
  "errors"
  "time"
  "math/rand"
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
)

func CreateMealHistories() {
  if DB.Migrator().HasTable(&models.MealHistory{}) {
    if err := DB.First(&models.MealHistory{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      morningTime := time.Date(2022, 12, 10, 7, 00, 00, 00, time.UTC)
      lunchTime := time.Date(2022, 12, 10, 12, 00, 00, 00, time.UTC)
      dinnerTime := time.Date(2022, 12, 10, 19, 00, 00, 00, time.UTC)
      snackTime := time.Date(2022, 12, 10, 15, 00, 00, 00, time.UTC)

      var man models.User
      DB.First(&man)

      var morningSession models.Session
      var lunchSession models.Session
      var dinnerSession models.Session
      var snackSession models.Session

      DB.Where("name = ?", "Morning").First(&morningSession)
      DB.Where("name = ?", "Lunch").First(&lunchSession)
      DB.Where("name = ?", "Dinner").First(&dinnerSession)
      DB.Where("name = ?", "Snack").First(&snackSession)

      for i := 0; i < 7; i++ {
        var dish1 models.Dish
        var dish2 models.Dish
        var dish3 models.Dish
        var dish4 models.Dish

        DB.Where("id = ?", rand.Intn(15) + 1).First(&dish1)
        DB.Where("id = ?", rand.Intn(15) + 1).First(&dish2)
        DB.Where("id = ?", rand.Intn(15) + 1).First(&dish3)
        DB.Where("id = ?", rand.Intn(15) + 1).First(&dish4)

        var mealHistories = []models.MealHistory{
          {Date: morningTime.AddDate(0, 0, i), Session: morningSession, Dish: dish1, User: man},
          {Date: lunchTime.AddDate(0, 0, i), Session: lunchSession, Dish: dish2, User: man},
          {Date: dinnerTime.AddDate(0, 0, i), Session: dinnerSession, Dish: dish3, User: man},
          {Date: snackTime.AddDate(0, 0, i), Session: snackSession, Dish: dish4, User: man},
        }

        DB.Create(&mealHistories)
      }
    }
  }
}
