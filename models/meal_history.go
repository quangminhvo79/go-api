package models

import (
  "time"
  "gorm.io/gorm"
)

type MealHistory struct {
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

func MealHistoriesByUser(uID int) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Model(&MealHistory{}).Where("user_id = ?", uID)
  }
}

func MealHistoriesBySession(sID int) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Model(&MealHistory{}).Where("session_id = ?", sID)
  }
}

func MealHistoryByID(ID int) *gorm.DB {
  return DB.Model(&MealHistory{}).Where("id = ?", ID)
}

func MealHistories() *gorm.DB {
  return DB.Model(&MealHistory{}).Joins("User").Joins("Dish").Joins("Session")
}

