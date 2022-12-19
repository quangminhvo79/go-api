package models

import "github.com/quangminhvo79/go-api/authentication"

type UserExercise struct {
  ID             uint64         `json:"id" gorm:"primary_key"`
  UserID         uint64         `json:"user_id" gorm:"index:,unique,composite:user_exercises_id"`
  ExerciseID     uint64         `json:"exercise_id" gorm:"index:,unique,composite:user_exercises_id"`
}

type UserExerciseInput struct {
  UserID         uint64    `json:"user_id"`
  ExerciseID     uint64    `json:"exercise_id" binding:"required"`
}

func (ue *UserExercise) AssignAttributes(input UserExerciseInput) {
  ue.ExerciseID = input.ExerciseID
  ue.UserID = authentication.UserID
}
