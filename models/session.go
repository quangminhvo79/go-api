package models

type Session struct {
  ID     uint64   `json:"id" gorm:"primary_key"`
  Name   string   `json:"name"`
}

type CreateSessionInput struct {
  Name  string `json:"name" binding:"required"`
}

type UpdateSessionInput struct {
  Name  string `json:"name" binding:"required"`
}
