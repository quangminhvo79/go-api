package models

import "gorm.io/gorm"

type Dish struct {
  gorm.Model
  ID       uint    `json:"id" gorm:"primary_key"`
  Name     string  `json:"name"`
  Calories float32 `json:"calories"`
}

type CreateDishInput struct {
  Name     string  `json:"name" binding:"required"`
  Calories float32 `json:"calories" binding:"required"`
}

type UpdateDishInput struct {
  Name   string  `json:"name"`
  Calories float32 `json:"calories"`
}
