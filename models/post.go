package models

import (
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/authentication"
)

type Post struct {
  gorm.Model
  Title          string   `json:"title"`
  Description    string   `json:"description"`
  UserID         uint64   `json:"user_id"`
  User           User
}

type PostInput struct {
  Title        string `json:"title" binding:"required"`
  Description  string `json:"description" binding:"required"`
  UserID       string `json:"user_id"`
}


func (p *Post) AssignAttributes(input PostInput) {
  p.Title = input.Title
  p.Description = input.Description
  p.UserID = authentication.UserID
}
