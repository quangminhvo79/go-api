package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/quangminhvo79/go-api/middlewares"
  "github.com/quangminhvo79/go-api/controllers"
)

func InitRouter() *gin.Engine {
  r := gin.Default()

  publicAPI := r.Group("/api")
  {
    publicAPI.POST("/token", controllers.GenerateToken)
    publicAPI.POST("/users/register", controllers.CreateUser)

    publicAPI.GET("/dishes", controllers.FindDishes)
    publicAPI.GET("/dishes/:id_or_name", controllers.FindDishesBy)

    publicAPI.GET("/sessions", controllers.FindSessions)
    publicAPI.GET("/sessions/:id", controllers.FindSessionsBy)

    publicAPI.GET("/exercises", controllers.FindExercises)
    publicAPI.GET("/exercises/:id_or_name", controllers.FindExercisesBy)

    publicAPI.GET("/posts", controllers.FindPosts)
    publicAPI.GET("/posts/:id", controllers.FindPost)
  }

  privateAPI := r.Group("/api").Use(middlewares.Auth())
  {
    privateAPI.GET("/users", controllers.FindUsers)
    privateAPI.GET("/users/:id", controllers.FindUser)

    privateAPI.GET("/sessions/:id/meal_histories", controllers.FindMealHistoriesBySession)

    privateAPI.GET("/meal_histories", controllers.FindMealHistories)
    privateAPI.GET("/meal_histories/:id", controllers.FindMealHistory)
    privateAPI.POST("/meal_histories", controllers.CreateMealHistory)

    privateAPI.GET("/achievement_rate", controllers.AchievementRate)
    privateAPI.GET("/body_fat_percent_graph", controllers.BodyFatPercentageGraph)

    privateAPI.GET("/user_exercises", controllers.FindUserExercises)
    privateAPI.POST("/user_exercises", controllers.CreateUserExercises)

    privateAPI.GET("/body_records", controllers.FindBodyRecords)
    privateAPI.POST("/body_records", controllers.CreateBodyRecord)

    privateAPI.GET("/diaries", controllers.FindDiaries)
    privateAPI.GET("/diaries/:id", controllers.FindDiary)
    privateAPI.POST("/diaries", controllers.CreateDiary)

    privateExerciseHistory := r.Group("/api/diaries/:id").Use(middlewares.Auth())
    {
      privateExerciseHistory.GET("/exercise_histories", controllers.FindExerciseHistories)
      privateExerciseHistory.GET("/exercise_histories/:exercise_history_id", controllers.FindExerciseHistory)
      privateExerciseHistory.POST("/exercise_histories", controllers.CreateExerciseHistory)
      privateExerciseHistory.PATCH("/exercise_histories/:exercise_history_id", controllers.UpdateExerciseHistory)
      privateExerciseHistory.DELETE("/exercise_histories/:exercise_history_id", controllers.DeleteExerciseHistory)
    }

    privateAPI.POST("/posts", controllers.CreatePost)
  }

  privateUser := r.Group("/api").Use(middlewares.Auth()).Use(middlewares.HasUser())
  {
    privateUser.PATCH("/users/:id", controllers.UpdateUser)
    privateUser.DELETE("/users/:id", controllers.DeleteUser)
  }

  privateMealHistory := r.Group("/api").Use(middlewares.Auth()).Use(middlewares.HasMealHistory())
  {
    privateMealHistory.PATCH("/meal_histories/:id", controllers.UpdateMealHistory)
    privateMealHistory.DELETE("/meal_histories/:id", controllers.DeleteMealHistory)
  }

  privateUserExercise := r.Group("/api").Use(middlewares.Auth()).Use(middlewares.HasUserExercise())
  {
    privateUserExercise.PATCH("/user_exercises/:id", controllers.UpdateUserExercises)
    privateUserExercise.DELETE("/user_exercises/:id", controllers.DeleteUserExercises)
  }

  privateBodyRecord := r.Group("/api").Use(middlewares.Auth()).Use(middlewares.HasBodyRecord())
  {
    privateBodyRecord.PATCH("/body_records/:id", controllers.UpdateBodyRecord)
    privateBodyRecord.DELETE("/body_records/:id", controllers.DeleteBodyRecord)
  }

  privateDiary := r.Group("/api").Use(middlewares.Auth()).Use(middlewares.HasDiary())
  {
    privateDiary.PATCH("/diaries/:id", controllers.UpdateDiary)
    privateDiary.DELETE("/diaries/:id", controllers.DeleteDiary)
  }

  privatePost := r.Group("/api").Use(middlewares.Auth()).Use(middlewares.HasPost())
  {
    privatePost.PATCH("/posts/:id", controllers.UpdatePost)
    privatePost.DELETE("/posts/:id", controllers.DeletePost)
  }

  return r
}
