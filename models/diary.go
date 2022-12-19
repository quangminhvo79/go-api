package models

import (
  "time"
  "github.com/quangminhvo79/go-api/authentication"
)

type Diary struct {
  ID                      uint64              `json:"id" gorm:"primary_key"`
  Date                    time.Time           `json:"date"`
  Note                    string              `json:"note"`
  UserID                  uint64              `json:"user_id"`
  User                    User
  ExerciseHistories       []ExerciseHistory   `json:"exercise_histories"`
}

type DiaryInput struct {
  Date                 time.Time          `json:"date" binding:"required"`
  Note                 string             `json:"note"`
  UserID               uint64             `json:"user_id"`
}

func (dr *Diary) AssignAttributes(input DiaryInput) {
  dr.Date = input.Date
  dr.Note = input.Note
  dr.UserID = authentication.UserID
}
