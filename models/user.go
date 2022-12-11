package models

type User struct {
  ID     		              uint   `json:"id" gorm:"primary_key"`
  Email  		              string `json:"email"`
  Password 	              string `json:"password"`
  Name                    string `json:"name"`
  AchievementWeightFrom   float32 `json:"achievement_weight_from"`
  AchievementWeightTo     float32 `json:"achievement_weight_to"`
}

type CreateUserInput struct {
  Email                   string `json:"email" binding:"required"`
  Password                string `json:"password" binding:"required"`
  Name                    string `json:"name" binding:"required"`
  AchievementWeightFrom   float32 `json:"achievement_weight_from"`
  AchievementWeightTo     float32 `json:"achievement_weight_to"`
}

type UpdateUserInput struct {
  Email                   string `json:"email"`
  Password                string `json:"password"`
  Name                    string `json:"name"`
  AchievementWeightFrom   float32 `json:"achievement_weight_from"`
  AchievementWeightTo     float32 `json:"achievement_weight_to"`
}
