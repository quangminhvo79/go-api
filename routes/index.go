package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/quangminhvo79/go-api/controllers"
)

func InitRouter() *gin.Engine {
  r := gin.Default()

  authRoutes := r.Group("api/auth")
  {
    authRoutes.POST("/register", controllers.FindUsers)
  }

  r.GET("/users", controllers.FindUsers)
  r.GET("/users/:id", controllers.FindUser)
  r.POST("/users", controllers.CreateUser)
  r.PATCH("/users/:id", controllers.UpdateUser)
  r.DELETE("/users/:id", controllers.DeleteUser)

  r.GET( "/meal_histories/:user_id", controllers.FindMealHistories)
  r.GET( "/meal_histories/:user_id/sessions/:session_id", controllers.FindMealHistoriesBySession)
  r.POST("/meal_histories", controllers.CreateMealHistory)
  r.PATCH("/meal_histories/:id", controllers.UpdateMealHistory)
  r.DELETE("/meal_histories/:id", controllers.DeleteMealHistory)

  r.GET( "/meal_histories/:user_id/achievement_rate", controllers.AchievementRate)
  r.GET( "/meal_histories/:user_id/body_fat_percent_graph", controllers.BodyFatPercentageGraph)

  r.GET( "/dishes", controllers.FindDishes)
  r.GET( "/dishes/:name", controllers.FindDishesByName)

  r.GET( "/sessions", controllers.FindSessions)
  r.GET( "/sessions/:name", controllers.FindSessionsByName)

  return r
}
