package seed

import (
  "errors"
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
)

func CreateExercises() {
  if DB.Migrator().HasTable(&models.Exercise{}) {
    if err := DB.First(&models.Exercise{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      var exercises = []models.Exercise{
        {Name: "Push Up", CaloriesBurned: 23},
        {Name: "Pull Up", CaloriesBurned: 43},
        {Name: "Lunges", CaloriesBurned: 65},
        {Name: "Squats", CaloriesBurned: 87},
        {Name: "Dumbbell rows", CaloriesBurned: 123},

        {Name: "Run", CaloriesBurned: 114},
        {Name: "Swiming", CaloriesBurned: 162},
        {Name: "Deadlifts", CaloriesBurned: 53},
        {Name: "Burpees", CaloriesBurned: 74},
        {Name: "Side planks", CaloriesBurned: 83},

        {Name: "Plank", CaloriesBurned: 90},
        {Name: "Glute bridge", CaloriesBurned: 80},
        {Name: "Leg press", CaloriesBurned: 85},
        {Name: "Bear Crawl", CaloriesBurned: 60},
        {Name: "Depth Jumps", CaloriesBurned: 45},
      }

      DB.Create(&exercises)
    }
  }
}
