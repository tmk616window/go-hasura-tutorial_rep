package graph

import (
	"api/graph/generated"
	"api/graph/model"
	"api/graph/models"
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

// func (s *GraphTestSuite) TestGetPrice() {
// db := s.resolver.DB

// }

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
