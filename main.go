package main

import (
  "github.com/quangminhvo79/go-api/models"
  "github.com/quangminhvo79/go-api/routes"
)

func main() {
  models.ConnectDatabase()
  defer models.CloseDatabaseConnection()
  r := routes.InitRouter()
  r.Run()
}
