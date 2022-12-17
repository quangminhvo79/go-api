package models

import (
  "time"
  "gorm.io/gorm"
)

type MealHistory struct {
  gorm.Model
  ID     		       uint       `json:"id" gorm:"primary_key"`
  Date  		       time.Time  `json:"date"`
  SessionID 	     uint       `json:"session_id"`
  DishID           uint       `json:"dish_id"`
  UserID           uint       `json:"user_id"`
  Session          Session
  Dish             Dish
  User             User
}

type CreateMealHistoryInput struct {
  Date          time.Time   `json:"date" binding:"required"`
  SessionID     uint        `json:"session_id" binding:"required"`
  DishID        uint        `json:"dish_id" binding:"required"`
  UserID        uint        `json:"user_id" binding:"required"`
}

type UpdateMealHistoryInput struct {
  Date          time.Time   `json:"date"`
  SessionID     uint        `json:"session_id"`
  DishID        uint        `json:"dish_id"`
  UserID        uint        `json:"user_id"`
}

type BodyFatPercentGraph struct {
  Date           time.Time  `json:"date"`
  TotalCalories  float32    `json:"total_calories"`
}
