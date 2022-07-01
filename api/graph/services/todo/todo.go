package todo

import (
	"api/graph/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ValidateTodoType struct {
	Title       string
	Description string
	LabelIDs    []int
	FinishTime  time.Time
	LabelCount  int
}

func CreateTodoLabel(labelIDs []int, todoID int, db *gorm.DB) {
	for _, labelID := range labelIDs {
		db.Create(&models.TodoLabel{
			LabelID: labelID,
			TodoID:  todoID,
		})
	}
}

func checkFinishTime(finishTime time.Time) time.Duration {
	now := time.Now()
	diff := now.Sub(finishTime)
	return diff
}

func ChangeTypeStringToTypeTime(stringFinishTime string) (time.Time, error) {
	layout := "2006-01-02 15:04"

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	finishTimeUTC, err := time.Parse(layout, stringFinishTime)
	if err != nil {
		return time.Time{}, err
	}

	finishTime := finishTimeUTC.In(jst)

	return finishTime, nil
}

func ValidateTodo(obj ValidateTodoType) error {
	if len(obj.Title) > 50 {
		return errors.New("タイトルは50文字にしてください")
	}

	if len(obj.Description) > 300 {
		return errors.New("説明を300文字以下にしてください")
	}

	diff := checkFinishTime(obj.FinishTime)
	if diff > 0 {
		return errors.New("終了期限を現在日時以降にしてください")
	}

	if len(obj.LabelIDs) > 5 || (5-obj.LabelCount)-len(obj.LabelIDs) < 0 {
		return errors.New("labelは登録できるのは5つまでです。")
	}
	return nil
}
