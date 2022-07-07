package graph

import (
	"api/graph/generated"
	"api/graph/model"
	"api/graph/models"
	createTodoService "api/graph/services/todo/create"
	"api/graph/test/util"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GraphTestSuite struct {
	util.Suite
	mutationResolver generated.MutationResolver
	queryResolver    generated.QueryResolver
	todoResolver     generated.TodoResolver
	todoLabel        generated.TodoLabelResolver
	resolver         Resolver
}

func (s *GraphTestSuite) SetupTest() {
	tx := s.DB().Begin()

	resolver := NewResolver(
		tx,
	)
	s.resolver = *resolver
	s.mutationResolver = s.resolver.Mutation()
	s.queryResolver = s.resolver.Query()
	s.todoResolver = s.resolver.Todo()
	s.todoLabel = s.resolver.TodoLabel()
}

func TestMain(t *testing.T) {
	suite.Run(t, new(GraphTestSuite))
}

// 各テストメソッドの実行後
func (s *GraphTestSuite) TearDownTest() {
	s.resolver.DB.Rollback()
	s.CloseDB()
}

// スイートの実行前
func (s *GraphTestSuite) SetupSuite() {
	s.SetupDB()
}

func (s *GraphTestSuite) TestGetTodo() {
	db := s.resolver.DB
	s.Run("正常系", func() {
		var sortInput *model.SortTodo
		var searchInput *model.SearchTodo

		result, _ := s.queryResolver.GqlgenTodos(context.Background(), sortInput, searchInput)

		var todos []*models.Todo
		db.Find(&todos)
		assert.Equal(s.T(), result, todos)
	})
}

func (s *GraphTestSuite) TestCreateTodo() {
	db := s.resolver.DB
	s.Run("正常系", func() {
		title := "testTitle"
		description := "testDescription"
		userID := 1
		priorityID := 1
		finishedTimeString := "2024-01-02 15:04"

		object := model.NewTodo{
			Title:       title,
			Description: description,
			UserID:      userID,
			PriorityID:  priorityID,
			FinishedAt:  finishedTimeString,
		}
		result, _ := s.mutationResolver.CreateTodo(context.Background(), object)

		var todo models.Todo
		db.Last(&todo)

		assert.Equal(s.T(), result.Title, todo.Title)
		assert.Equal(s.T(), result.Description, todo.Description)
		assert.Equal(s.T(), result.PriorityID, todo.PriorityID)
		assert.Equal(s.T(), result.StatusID, todo.StatusID)

	})

	s.Run("異常系", func() {
		s.Run("タイトルが50字以上であればエラーを返す", func() {
			title := "testTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitle"
			description := "testDescription"
			userID := 1
			priorityID := 1
			finishedTimeString := "2024-01-02 15:04"

			object := model.NewTodo{
				Title:       title,
				Description: description,
				UserID:      userID,
				PriorityID:  priorityID,
				FinishedAt:  finishedTimeString,
			}
			_, err := s.mutationResolver.CreateTodo(context.Background(), object)
			assert.Equal(s.T(), err.Error(), "タイトルは50文字以下にしてください")

		})

		s.Run("説明文が300字以上", func() {
			title := "testTitle"
			description := "testDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescription"
			userID := 1
			priorityID := 1
			finishedTimeString := "2024-01-02 15:04"

			object := model.NewTodo{
				Title:       title,
				Description: description,
				UserID:      userID,
				PriorityID:  priorityID,
				FinishedAt:  finishedTimeString,
			}
			_, err := s.mutationResolver.CreateTodo(context.Background(), object)
			assert.Equal(s.T(), err.Error(), "説明を300文字以下にしてください")
		})

		s.Run("ラベルの登録が5つ以上の時にエラーが発生する", func() {
			title := "testTitle"
			description := "testDescription"
			userID := 1
			priorityID := 1
			finishedTimeString := "2024-01-02 15:04"
			labelIDs := []int{1, 2, 3, 4, 5, 6}

			object := model.NewTodo{
				Title:       title,
				Description: description,
				UserID:      userID,
				LabelIDs:    labelIDs,
				PriorityID:  priorityID,
				FinishedAt:  finishedTimeString,
			}
			_, err := s.mutationResolver.CreateTodo(context.Background(), object)
			assert.Equal(s.T(), err.Error(), "labelは登録できるのは5つまでです。")
		})

		s.Run("終了時間は現在時刻以前にするとエラーが発生する", func() {
			title := "testTitle"
			description := "testDescription"
			userID := 1
			priorityID := 1
			finishedTimeString := "2020-01-02 15:04"

			object := model.NewTodo{
				Title:       title,
				Description: description,
				UserID:      userID,
				PriorityID:  priorityID,
				FinishedAt:  finishedTimeString,
			}
			_, err := s.mutationResolver.CreateTodo(context.Background(), object)
			assert.Equal(s.T(), err.Error(), "終了期限を現在日時以降にしてください")
		})
	})
}

func (s *GraphTestSuite) TestUpdateTodo() {
	db := s.resolver.DB
	s.Run("正常系", func() {
		todo, _ := createTodoService.CreateTodo(db, model.NewTodo{
			Title:       "testTitle",
			Description: "testDescription",
			UserID:      1,
			PriorityID:  1,
			FinishedAt:  "2024-01-02 15:04",
		})
		id := todo.ID
		title := "updateTitle"
		description := "updateDescription"
		finishTimeString := "2025-01-02 15:04"

		object := model.UpdateTodo{
			ID:          id,
			Title:       title,
			Description: description,
			FinishedAt:  finishTimeString,
		}
		result, _ := s.mutationResolver.UpdateTodo(context.Background(), object)

		assert.Equal(s.T(), result.ID, id)
		assert.Equal(s.T(), result.Title, title)
		assert.Equal(s.T(), result.Description, description)
	})

	s.Run("異常系", func() {
		s.Run("タイトルが50字以上であればエラーを返す", func() {

			todo, _ := createTodoService.CreateTodo(db, model.NewTodo{
				Title:       "testTitle",
				Description: "testDescription",
				UserID:      1,
				PriorityID:  1,
				FinishedAt:  "2024-01-02 15:04",
			})

			id := todo.ID
			title := "testTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitle"

			object := model.UpdateTodo{
				ID:    id,
				Title: title,
			}
			_, err := s.mutationResolver.UpdateTodo(context.Background(), object)

			assert.Equal(s.T(), err.Error(), "タイトルは50文字以下にしてください")
		})

		s.Run("説明文が300字以上であればエラーを返す", func() {
			todo, _ := createTodoService.CreateTodo(db, model.NewTodo{
				Title:       "testTitle",
				Description: "testDescription",
				UserID:      1,
				PriorityID:  1,
				FinishedAt:  "2024-01-02 15:04",
			})

			id := todo.ID
			description := "testDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescriptiontestDescription"

			object := model.UpdateTodo{
				ID:          id,
				Description: description,
			}
			_, err := s.mutationResolver.UpdateTodo(context.Background(), object)

			assert.Equal(s.T(), err.Error(), "説明を300文字以下にしてください")
		})

		s.Run("ラベルの登録が5つ以上の時にエラーが発生する", func() {
			todo, _ := s.mutationResolver.CreateTodo(context.Background(), model.NewTodo{
				Title:       "testTitle",
				Description: "testDescription",
				UserID:      1,
				PriorityID:  1,
				LabelIDs:    []int{1, 2, 3},
				FinishedAt:  "2024-01-02 15:04",
			})

			id := todo.ID
			addLabelIDs := []int{4, 5, 6}

			object := model.UpdateTodo{
				ID:          id,
				Title:       "updateTitle",
				Description: "updateDescription",
				AddLabelIDs: addLabelIDs,
				FinishedAt:  "2025-01-02 15:04",
			}
			_, err := s.mutationResolver.UpdateTodo(context.Background(), object)

			assert.Equal(s.T(), err.Error(), "labelは登録できるのは5つまでです。")
		})

		s.Run("終了時間は現在時刻以前にするとエラーが発生する", func() {
			todo, _ := createTodoService.CreateTodo(db, model.NewTodo{
				Title:       "testTitle",
				Description: "testDescription",
				UserID:      1,
				PriorityID:  1,
				FinishedAt:  "2024-01-02 15:04",
			})

			id := todo.ID
			finishTimeString := "2020-01-02 15:04"

			object := model.UpdateTodo{
				ID:         id,
				FinishedAt: finishTimeString,
			}
			_, err := s.mutationResolver.UpdateTodo(context.Background(), object)

			assert.Equal(s.T(), err.Error(), "終了期限を現在日時以降にしてください")
		})
	})
}
