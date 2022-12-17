package main

import (
  "github.com/quangminhvo79/go-api/database"
  "github.com/quangminhvo79/go-api/routes"
)

func main() {
  database.ConnectDatabase()
  database.Migrate()

  defer database.CloseDatabaseConnection()

  r := routes.InitRouter()
  r.Run()
}
