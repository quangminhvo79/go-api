package models

type Exercise struct {
  ID       			 uint64    `json:"id" gorm:"primary_key"`
  Name     			 string  `json:"name"`
  CaloriesBurned float32 `json:"calories_burned"`
}

type CreateExerciseInput struct {
  Name     			 string  `json:"name" binding:"required"`
  CaloriesBurned float32 `json:"calories_burned" binding:"required"`
}

type UpdateExerciseInput struct {
  Name   				 string  `json:"name"`
  CaloriesBurned float32 `json:"calories_burned"`
}
