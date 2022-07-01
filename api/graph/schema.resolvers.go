package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"api/graph/models"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTodoLabel(ctx context.Context, input model.NewTodo) (*models.TodoLabel, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*models.Todo, error) {
	db := r.Resolver.DB
	var todo models.Todo
	err := db.Delete(todo, input.ID).Error
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
