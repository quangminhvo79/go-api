package models

import (
  "gorm.io/gorm"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  gorm.Model
  ID     		              uint   `json:"id" gorm:"primary_key"`
  Email  		              string `json:"email"`
  Password 	              string `json:"password"`
  Username                string `json:"username"`
  AchievementWeightFrom   float32 `json:"achievement_weight_from"`
  AchievementWeightTo     float32 `json:"achievement_weight_to"`
}

type UserOutput struct {
  ID                      uint   `json:"id"`
  Email                   string `json:"email"`
  Username                string `json:"username"`
  AchievementWeightFrom   float32 `json:"achievement_weight_from"`
  AchievementWeightTo     float32 `json:"achievement_weight_to"`
}

type CreateUserInput struct {
  Email                   string `json:"email" binding:"required"`
  Password                string `json:"password" binding:"required"`
  Username                string `json:"username" binding:"required"`
  AchievementWeightFrom   float32 `json:"achievement_weight_from"`
  AchievementWeightTo     float32 `json:"achievement_weight_to"`
}

type UpdateUserInput struct {
  Email                   string `json:"email"`
  Password                string `json:"password"`
  Username                string `json:"username"`
  AchievementWeightFrom   float32 `json:"achievement_weight_from"`
  AchievementWeightTo     float32 `json:"achievement_weight_to"`
}

func (user *User) HashPassword(password string) error {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  if err != nil {
    return err
  }
  user.Password = string(bytes)
  return nil
}

func (user *User) CheckPassword(providedPassword string) error {
  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
  if err != nil {
    return err
  }
  return nil
}

func (user *User) UserResponseData() UserOutput {
  return UserOutput{
    ID: user.ID,
    Email: user.Email,
    Username: user.Username,
    AchievementWeightFrom: user.AchievementWeightFrom,
    AchievementWeightTo: user.AchievementWeightTo,
  }
}
