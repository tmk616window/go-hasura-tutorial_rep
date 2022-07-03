package todoLabel

import (
	"api/graph/models"

	"gorm.io/gorm"
)

func CreateTodoLabel(db *gorm.DB, labelID int, todoID int) error {
	err := db.Create(&models.TodoLabel{
		LabelID: labelID,
		TodoID:  todoID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
