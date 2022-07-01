package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"api/graph/models"
	servicesTodo "api/graph/services/todo"
	"context"
	"fmt"
	"time"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	db := r.Resolver.DB
	const defaultStatus = 1

	// フォーマット　"2022/6/28 13:00"
	finishTime, err := servicesTodo.ChangeTypeStringToTypeTime(input.FinishedAt)
	if err != nil {
		return nil, err
	}

	err = servicesTodo.ValidateTodo(servicesTodo.ValidateTodoType{
		Title:       input.Title,
		Description: input.Description,
		LabelIDs:    input.LabelIDs,
		FinishTime:  finishTime,
		LabelCount:  0,
	})
	if err != nil {
		return nil, err
	}

	t := &models.Todo{
		Title:       input.Title,
		Description: input.Description,
		UserID:      input.PriorityID,
		StatusID:    defaultStatus,
		PriorityID:  input.PriorityID,
		FinishedAt:  finishTime,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = db.Create(t).Error
	if err != nil {
		return nil, err
	}

	for _, labelID := range input.LabelIDs {
		db.Create(&models.TodoLabel{
			LabelID: labelID,
			TodoID:  t.ID,
		})
	}

	return t, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*models.Todo, error) {
	db := r.Resolver.DB
	todo := models.Todo{}

	finishTime, err := servicesTodo.ChangeTypeStringToTypeTime(input.FinishedAt)
	if err != nil {
		return nil, err
	}

	var todoLabel []*models.TodoLabel
	var labelCount int64

	db.
		Where("todo_id = ?", input.ID).
		Find(&todoLabel).
		Count(&labelCount)

	validateTodo := servicesTodo.ValidateTodoType{
		Title:       input.Title,
		Description: input.Description,
		LabelIDs:    input.LabelIDs,
		FinishTime:  finishTime,
		LabelCount:  int(labelCount),
	}

	err = servicesTodo.ValidateTodo(validateTodo)
	if err != nil {
		return nil, err
	}

	if len(input.LabelIDs) != 0 {
		for _, labelID := range input.LabelIDs {
			db.Create(&models.TodoLabel{
				LabelID: labelID,
				TodoID:  input.ID,
			})
		}
	}

	err = db.Model(&todo).First(&todo, input.ID).Updates(models.Todo{
		Title:       input.Title,
		Description: input.Description,
	}).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *queryResolver) GqlgenTodos(ctx context.Context, sortInput *model.SortTodo, searchInput *model.SearchTodo) ([]*models.Todo, error) {
	var todos []*models.Todo
	db := r.Resolver.DB

	if searchInput != nil {
		db.Where(searchInput.Column+" = ?", searchInput.Value)
	} else if sortInput != nil {
		db.Order(sortInput.Column + " " + string(sortInput.Sort))
	}
	db.Find(&todos)

	return todos, nil
}

func (r *todoResolver) Status(ctx context.Context, obj *models.Todo) (*model.Status, error) {
	var status model.Status
	db := r.Resolver.DB
	db.First(&status, obj.StatusID)
	return &status, nil
}

func (r *todoResolver) Priority(ctx context.Context, obj *models.Todo) (*model.Priority, error) {
	var priority model.Priority
	db := r.Resolver.DB
	db.First(&priority, obj.PriorityID)
	return &priority, nil
}

func (r *todoResolver) TodoLabels(ctx context.Context, obj *models.Todo) ([]*models.TodoLabel, error) {
	var todoLabel []*models.TodoLabel
	db := r.Resolver.DB
	db.Where("todo_id = ?", obj.ID).Find(&todoLabel)
	return todoLabel, nil
}

func (r *todoLabelResolver) Label(ctx context.Context, obj *models.TodoLabel) (*model.Label, error) {
	var label model.Label
	db := r.Resolver.DB
	db.First(&label, obj.LabelID)
	return &label, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// TodoLabel returns generated.TodoLabelResolver implementation.
func (r *Resolver) TodoLabel() generated.TodoLabelResolver { return &todoLabelResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type todoLabelResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateTodoLabel(ctx context.Context, input model.NewTodo) (*models.TodoLabel, error) {
	panic(fmt.Errorf("not implemented"))
}
