package models

import (
  "time"
  "github.com/quangminhvo79/go-api/authentication"
)

type MealHistory struct {
  ID     		       uint64       `json:"id" gorm:"primary_key"`
  Date  		       time.Time  `json:"date"`
  SessionID 	     uint64       `json:"session_id"`
  DishID           uint64       `json:"dish_id"`
  UserID           uint64       `json:"user_id"`
  Session          Session
  Dish             Dish
  User             User
}

type MealHistoryInput struct {
  Date          time.Time   `json:"date" binding:"required"`
  SessionID     uint64        `json:"session_id" binding:"required"`
  DishID        uint64        `json:"dish_id" binding:"required"`
  UserID        uint64        `json:"user_id"`
}

type BodyFatPercentGraph struct {
  Date           time.Time  `json:"date"`
  TotalCalories  float32    `json:"total_calories"`
}

func (mh *MealHistory) AssignAttributes(input MealHistoryInput) {
  mh.Date = input.Date
  mh.SessionID = input.SessionID
  mh.DishID = input.DishID
  mh.UserID = authentication.UserID
}
