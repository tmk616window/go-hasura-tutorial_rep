package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"api/graph/models"
	postgresql "api/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTodoLabel(ctx context.Context, input model.NewTodo) (*models.TodoLabel, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, sortInput *model.SortTodo, searchInput *model.SearchTodo) ([]*models.Todo, error) {
	var todos []*models.Todo
	db := postgresql.DbConnect()
	if (sortInput) != nil && (searchInput) != nil {
		db.Preload("TodoLabels").Find(&todos)
	} else if searchInput != nil {
		db.Preload("TodoLabels").Where(searchInput.Column+" = ?", searchInput.Value).Find(&todos)
	} else if sortInput != nil {
		db.Preload("TodoLabels").Order(sortInput.Column + " " + sortInput.Value).Find(&todos)
	}
	fmt.Println(&todos)
	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Status(ctx context.Context, obj *models.Todo) (*model.Status, error) {
	var status model.Status
	db := postgresql.DbConnect()
	db.First(&status, obj.StatusID)
	fmt.Println(status)
	return &status, nil
}

func (r *todoResolver) Priority(ctx context.Context, obj *models.Todo) (*model.Priority, error) {
	var priority model.Priority
	db := postgresql.DbConnect()
	db.First(&priority, obj.PriorityID)
	fmt.Println(priority)
	return &priority, nil
}

func (r *todoResolver) TodoLabels(ctx context.Context, obj *models.Todo) ([]*models.TodoLabel, error) {
	var todoLabel []*models.TodoLabel
	db := postgresql.DbConnect()
	db.Preload("Labels").Find(&todoLabel)
	return todoLabel, nil
}

func (r *todoLabelResolver) Label(ctx context.Context, obj *models.TodoLabel) (*model.Label, error) {
	var label model.Label
	db := postgresql.DbConnect()
	db.First(&label)
	fmt.Println(label)
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
func (r *queryResolver) Users(ctx context.Context, column string, value string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) SortTodos(ctx context.Context, column string, value string) ([]*models.Todo, error) {
	var sortTodos []*models.Todo
	db := postgresql.DbConnect()
	db.Order(column + " " + value).Find(&sortTodos)
	fmt.Println(&sortTodos)
	return sortTodos, nil
}
func (r *todoResolver) Labels(ctx context.Context, obj *models.Todo) ([]*model.Label, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *todoResolver) PriorityID(ctx context.Context, obj *models.Todo) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *todoResolver) StatusID(ctx context.Context, obj *models.Todo) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
