package database

import (
  "gorm.io/gorm"
  sqlite "gorm.io/driver/sqlite"
)

var DB *gorm.DB;

func ConnectDatabase() {
  database, err := gorm.Open(sqlite.Open("arent_minhvo.db"), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

  DB = database
}

func CloseDatabaseConnection() {
  dbSQL, err := DB.DB()
  if err != nil {
    panic("Failed to close connection from database")
  }
  dbSQL.Close()
}
