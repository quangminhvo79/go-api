package models

import "gorm.io/gorm"

type Session struct {
  gorm.Model
  ID     uint   `json:"id" gorm:"primary_key"`
  Name   string `json:"name"`
}

type CreateSessionInput struct {
  Name  string `json:"name" binding:"required"`
}

type UpdateSessionInput struct {
  Name  string `json:"name" binding:"required"`
}
