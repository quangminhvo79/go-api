package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/quangminhvo79/go-api/middlewares"
  "github.com/quangminhvo79/go-api/controllers"
)

func InitRouter() *gin.Engine {
  r := gin.Default()

  publicAPI := r.Group("/api")
  publicAPI.POST("/token", controllers.GenerateToken)
  publicAPI.POST("/users/register", controllers.CreateUser)

  privateAPI := r.Group("/api").Use(middlewares.Auth())
  privateAPI.GET("/users", controllers.FindUsers)
  privateAPI.GET("/users/:id", controllers.FindUser)
  privateAPI.PATCH("/users/:id", controllers.UpdateUser)
  privateAPI.DELETE("/users/:id", controllers.DeleteUser)

  privateAPI.GET( "/meal_histories/:user_id", controllers.FindMealHistories)
  privateAPI.GET( "/meal_histories/:user_id/sessions/:session_id", controllers.FindMealHistoriesBySession)
  privateAPI.POST("/meal_histories", controllers.CreateMealHistory)
  privateAPI.PATCH("/meal_histories/:id", controllers.UpdateMealHistory)
  privateAPI.DELETE("/meal_histories/:id", controllers.DeleteMealHistory)

  privateAPI.GET( "/meal_histories/:user_id/achievement_rate", controllers.AchievementRate)
  privateAPI.GET( "/meal_histories/:user_id/body_fat_percent_graph", controllers.BodyFatPercentageGraph)

  privateAPI.GET( "/dishes", controllers.FindDishes)
  privateAPI.GET( "/dishes/:id_or_name", controllers.FindDishesByName)

  privateAPI.GET( "/sessions", controllers.FindSessions)
  privateAPI.GET( "/sessions/:id_or_name", controllers.FindSessionsByName)

  privateAPI.GET( "/exercises", controllers.FindExercises)
  privateAPI.GET( "/exercises/:id_or_name", controllers.FindExercisesBy)


  return r
}
