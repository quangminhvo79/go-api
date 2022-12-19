package models

type Dish struct {
  ID       uint64    `json:"id" gorm:"primary_key"`
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
