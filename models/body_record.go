package models

import (
  "time"
  "github.com/quangminhvo79/go-api/authentication"
)

type BodyRecord struct {
  ID       uint64         `json:"id" gorm:"primary_key"`
  Date     time.Time      `json:"date"`
  Weight   float32        `json:"weight"`
  Height   float32        `json:"height"`
  UserID   uint64         `json:"user_id"`
  User     User
}

type BodyRecordInput struct {
  Date     time.Time      `json:"date" binding:"required"`
  Weight   float32        `json:"weight" binding:"required"`
  Height   float32        `json:"height" binding:"required"`
}

func (br *BodyRecord) AssignAttributes(input BodyRecordInput) {
  br.Date = input.Date
  br.Weight = input.Weight
  br.Height = input.Height
  br.UserID = authentication.UserID
}
