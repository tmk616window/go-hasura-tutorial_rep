package todo

import (
	"api/graph/models"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ValidateTodoType struct {
	Title       string
	Description string
	LabelList   string
	FinishTime  time.Time
	LabelCount  int
}

func CreateTodoLabelRelation(labels string, todoID int, db *gorm.DB) {
	labelArr := strings.Split(labels, " ")

	labelList := map[string]int{
		"label1": 1,
		"label2": 2,
		"label3": 3,
		"label4": 4,
		"label5": 5,
		"label6": 6,
		"label7": 7,
		"label8": 8,
		"label9": 9,
	}

	for _, label := range labelArr {
		labelID := labelList[label]
		db.Create(&models.TodoLabel{
			LabelID: labelID,
			TodoID:  todoID,
		})
	}
}

func compareTime(finishTime time.Time) time.Duration {
	now := time.Now()
	diff := now.Sub(finishTime)
	return diff
}

func StringToTime(stringFinishTime string) time.Time {
	layout := "2006-01-02 15:04"

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	finishTimeUTC, _ := time.Parse(layout, stringFinishTime)
	finishTime := finishTimeUTC.In(jst)

	return finishTime
}

func TodoValidate(obj ValidateTodoType) error {
	if len(obj.Title) > 50 {
		return errors.New("タイトルは50文字にしてください")
	}

	if len(obj.Description) > 300 {
		return errors.New("説明を300文字以下にしてください")
	}

	diff := compareTime(obj.FinishTime)
	if diff > 0 {
		return errors.New("終了期限を現在日時以降にしてください")
	}

	labelArr := strings.Split(obj.LabelList, " ")
	if len(labelArr) > 5 || (5-obj.LabelCount)-len(labelArr) < 0 {
		return errors.New("labelは登録できるのは5つまでです。")
	}

	return nil
}
