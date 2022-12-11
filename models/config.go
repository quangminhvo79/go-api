package models

import (
  "errors"
  "time"
  "math/rand"
  "gorm.io/gorm"
  sqlite "gorm.io/driver/sqlite"
)

var DB *gorm.DB;

func ConnectDatabase() {
  database, err := gorm.Open(sqlite.Open("arent_minhvo.db"), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

  err = database.AutoMigrate(&User{}, &Dish{}, &MealHistory{}, &Session{})
  if err != nil {
    return
  }

  seeds(database)
  DB = database
}

func CloseDatabaseConnection() {
  dbSQL, err := DB.DB()
  if err != nil {
    panic("Failed to close connection from database")
  }
  dbSQL.Close()
}

func seeds(db *gorm.DB) {

  if db.Migrator().HasTable(&User{}) {
    if err := db.First(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      var users = []User{
        {Email: "man@gmail.com", Password: "12345", Name: "Man", AchievementWeightFrom: 60, AchievementWeightTo: 75},
        {Email: "woman@gmail.com", Password: "12345", Name: "Woman", AchievementWeightFrom: 60, AchievementWeightTo: 49},
      }

      db.Create(&users)
    }
  }

  if db.Migrator().HasTable(&Dish{}) {
    if err := db.First(&Dish{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      var dishes = []Dish{
        {Name: "Beefsteak", Calories: 180},
        {Name: "Seafood soup", Calories: 150},
        {Name: "Chicken", Calories: 160},
        {Name: "Fried rice", Calories: 110},
        {Name: "Shrimp in batter", Calories: 145},

        {Name: "Big Beefsteak", Calories: 390},
        {Name: "Big Seafood soup", Calories: 310},
        {Name: "Big Chicken", Calories: 330},
        {Name: "Big Fried rice", Calories: 300},
        {Name: "Big Shrimp in batter", Calories: 375},

        {Name: "Small Beefsteak", Calories: 90},
        {Name: "Small Seafood soup", Calories: 80},
        {Name: "Small Chicken", Calories: 85},
        {Name: "Small Fried rice", Calories: 60},
        {Name: "Small Shrimp in batter", Calories: 45},
      }

      db.Create(&dishes)
    }
  }

  if db.Migrator().HasTable(&Session{}) {
    if err := db.First(&Session{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      var sessions = []Session{{Name: "Morning"}, {Name: "Lunch"}, {Name: "Dinner"}, {Name: "Snack"}}
      db.Create(&sessions)
    }
  }

  if db.Migrator().HasTable(&MealHistory{}) {
    if err := db.First(&MealHistory{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      morningTime := time.Date(2022, 12, 10, 7, 00, 00, 00, time.UTC)
      lunchTime := time.Date(2022, 12, 10, 12, 00, 00, 00, time.UTC)
      dinnerTime := time.Date(2022, 12, 10, 19, 00, 00, 00, time.UTC)
      snackTime := time.Date(2022, 12, 10, 15, 00, 00, 00, time.UTC)

      var man User
      db.First(&man)

      var morningSession Session
      var lunchSession Session
      var dinnerSession Session
      var snackSession Session

      db.Where("name = ?", "Morning").First(&morningSession)
      db.Where("name = ?", "Lunch").First(&lunchSession)
      db.Where("name = ?", "Dinner").First(&dinnerSession)
      db.Where("name = ?", "Snack").First(&snackSession)

      for i := 0; i < 7; i++ {
        var dish1 Dish
        var dish2 Dish
        var dish3 Dish
        var dish4 Dish

        db.Where("id = ?", rand.Intn(15) + 1).First(&dish1)
        db.Where("id = ?", rand.Intn(15) + 1).First(&dish2)
        db.Where("id = ?", rand.Intn(15) + 1).First(&dish3)
        db.Where("id = ?", rand.Intn(15) + 1).First(&dish4)

        var mealHistories = []MealHistory{
          {Date: morningTime.AddDate(0, 0, i), Session: morningSession, Dish: dish1, User: man},
          {Date: lunchTime.AddDate(0, 0, i), Session: lunchSession, Dish: dish2, User: man},
          {Date: dinnerTime.AddDate(0, 0, i), Session: dinnerSession, Dish: dish3, User: man},
          {Date: snackTime.AddDate(0, 0, i), Session: snackSession, Dish: dish4, User: man},
        }

        db.Create(&mealHistories)
      }
    }
  }
}
