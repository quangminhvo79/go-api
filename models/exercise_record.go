package models

import (
  "time"
  "gorm.io/gorm"
)

type ExerciseRecord struct {
  gorm.Model
  ID     		       uint       `json:"id" gorm:"primary_key"`
  Date  		       time.Time  `json:"date"`
  ExerciseID       uint       `json:"exercise_id"`
  UserID           uint       `json:"user_id"`
  Exercise         Exercise
  User             User
}
