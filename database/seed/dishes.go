package seed

import (
  "errors"
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
)

func CreateDishes() {
  if DB.Migrator().HasTable(&models.Dish{}) {
    if err := DB.First(&models.Dish{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      var dishes = []models.Dish{
        {Name: "Beefsteak", Calories: 180},
        {Name: "Seafood soup", Calories: 150},
        {Name: "Chicken", Calories: 160},
        {Name: "Fried rice", Calories: 110},
        {Name: "Shrimp in batter", Calories: 145},

        {Name: "Big Beefsteak", Calories: 390},
        {Name: "Big Seafood soup", Calories: 310},
        {Name: "Big Chicken", Calories: 330},
        {Name: "Big Fried rice", Calories: 300},
        {Name: "Big Shrimp in batter", Calories: 375},

        {Name: "Small Beefsteak", Calories: 90},
        {Name: "Small Seafood soup", Calories: 80},
        {Name: "Small Chicken", Calories: 85},
        {Name: "Small Fried rice", Calories: 60},
        {Name: "Small Shrimp in batter", Calories: 45},
      }

      DB.Create(&dishes)
    }
  }
}
