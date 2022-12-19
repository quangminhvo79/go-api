package seed

import (
  "errors"
  "time"
  "math/rand"
  "fmt"
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
)

func CreateDiary() {
  // DB.Migrator().DropTable(&models.Diary{})
  if DB.Migrator().HasTable(&models.Diary{}) {
    if err := DB.First(&models.Diary{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
      date := time.Date(2022, 12, 10, 7, 00, 00, 00, time.UTC)

      var man models.User
      DB.First(&man)

      for i := 0; i < 7; i++ {
        var exercisesHistories []models.ExerciseHistory

        for ei := 0; ei < 4; ei++ {
          exercisesHistories = append(exercisesHistories, models.ExerciseHistory{ExerciseID: uint64(rand.Intn(15) + 1)})
        }

        var diaries = []models.Diary{
          {Date: date.AddDate(0, 0, i), Note: fmt.Sprintf("Diary Note %d", i), User: man, ExerciseHistories: exercisesHistories},
        }

        DB.Create(&diaries)
      }
    }
  }
}
