package models

type ExerciseHistory struct {
  ID     		       uint64       `json:"id" gorm:"primary_key"`
  DiaryID          uint64       `json:"diary_id"`
  ExerciseID       uint64       `json:"exercise_id"`
  Exercise         Exercise
}

type ExerciseHistoryInput struct {
  DiaryID       uint64        `json:"diary_id" binding:"required"`
  ExerciseID    uint64        `json:"exercise_id" binding:"required"`
}

func (eh *ExerciseHistory) AssignAttributes(input ExerciseHistoryInput) {
  eh.DiaryID = input.DiaryID
  eh.ExerciseID = input.ExerciseID
}
