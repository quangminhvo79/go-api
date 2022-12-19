package database

import (
  "github.com/quangminhvo79/go-api/models"
  "github.com/quangminhvo79/go-api/database/seed"
)

func Migrate() {
  DB.AutoMigrate(
    &models.User{},
    &models.Dish{},
    &models.MealHistory{},
    &models.Session{},
    &models.Exercise{},
    &models.UserExercise{},
    &models.Diary{},
    &models.ExerciseHistory{},
    &models.BodyRecord{},
    &models.Post{},
  )

  seed.DB = DB
  seed.CreateUsers()
  seed.CreateDishes()
  seed.CreateSessions()
  seed.CreateMealHistories()
  seed.CreateExercises()
  seed.CreateDiary()
}
