package graph

import (
	"api/graph/generated"
	"api/graph/models"
	"api/graph/services/common"
	"api/graph/test/util"
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
		statusID := 1
		finishedTimeTypeString := "2024-6-28 13:00"
		finishTimeTypeTime, _ := common.ChangeStringToTime(finishedTimeTypeString)

		object := &models.Todo{
			Title:       title,
			Description: description,
			UserID:      userID,
			StatusID:    statusID,
			PriorityID:  priorityID,
			FinishedAt:  finishTimeTypeTime,
		}
		db.Create(object)

		var todo models.Todo
		db.Last(&todo)

		assert.Equal(s.T(), object.Title, todo.Title)
		assert.Equal(s.T(), object.Description, todo.Description)
		assert.Equal(s.T(), object.PriorityID, todo.PriorityID)
		assert.Equal(s.T(), object.StatusID, todo.StatusID)
		// s.test.Fail()
	})

	s.Run("異常系", func() {
		s.Run("タイトルが50字以上であればエラーを返す", func() {
			title := "testTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitletestTitle"
			description := "testDescription"
			userID := 1
			priorityID := 1
			statusID := 1
			finishedTimeTypeString := "2024/6/28 13:00"
			finishTimeTypeTime, _ := common.ChangeStringToTime(finishedTimeTypeString)

			object := &models.Todo{
				Title:       title,
				Description: description,
				UserID:      userID,
				StatusID:    statusID,
				PriorityID:  priorityID,
				FinishedAt:  finishTimeTypeTime,
			}
			db.Create(object)

		})

		s.Run("説明文が300字以上", func() {

		})

		s.Run("ラベルの登録が5つ以上の時にエラーが発生する", func() {

		})

		s.Run("終了時間は現在時刻以前にするとエラーが発生する", func() {

		})
	})

}
