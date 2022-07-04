// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"api/graph/models"
	"fmt"
	"io"
	"strconv"
)

type DeleteTodo struct {
	ID int `json:"id"`
}

type Label struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LabelIDs    []int  `json:"labelIDs"`
	UserID      int    `json:"userID"`
	PriorityID  int    `json:"PriorityID"`
	FinishedAt  string `json:"finishedAt"`
}

type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SearchTodo struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}

type SortTodo struct {
	Column string `json:"column"`
	Sort   Sort   `json:"sort"`
}

type Status struct {
	ID    int            `json:"id"`
	Name  string         `json:"name"`
	Todos []*models.Todo `json:"todos"`
}

type User struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Todos     []*models.Todo `json:"todos"`
	SortTodos []*models.Todo `json:"sortTodos"`
}

type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

var AllSort = []Sort{
	SortAsc,
	SortDesc,
}

func (e Sort) IsValid() bool {
	switch e {
	case SortAsc, SortDesc:
		return true
	}
	return false
}

func (e Sort) String() string {
	return string(e)
}

func (e *Sort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Sort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Sort", str)
	}
	return nil
}

func (e Sort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
