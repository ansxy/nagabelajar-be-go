package test

import (
	"errors"
	"testing"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

// CategoryUsecaseTestSuite is a test suite for the Category use case
type CategoryUsecaseTestSuite struct {
	BaseTestSuite
	uc usecase.IFaceUsecase
}

// SetupTest sets up the test environment for Category use case
func (s *CategoryUsecaseTestSuite) SetupTest() {
	s.BaseTestSuite.SetupTest()
	s.uc = usecase.NewUsecase(&usecase.Usecase{
		Repo:      s.mockRepo,
		DB:        nil, // assuming DB is not needed for the usecase or provide actual db connection if needed
		Validator: s.validate,
	})
}

// TestCreateCategory tests the CreateCategory use case method
func (s *CategoryUsecaseTestSuite) TestCreateCategory() {
	// Valid input
	s.Run("Valid input", func() {
		req := &request.UpsertCategoryRequest{
			Name: "Test Category",
		}

		s.mockRepo.EXPECT().CreateCategory(s.ctx, gomock.Any()).Return(nil)

		err := s.uc.CreateCategory(s.ctx, req)

		assert.NoError(s.T(), err, "Expected no error while creating category")
	})

	// Missing required fields
	s.Run("Missing required fields", func() {
		req := &request.UpsertCategoryRequest{}

		// Validate struct
		err := s.validateStruct(req)
		assert.Error(s.T(), err, "Expected error due to missing required fields")

		// No need to expect repository calls as validation should fail
	})
}

// TestDeleteOneCategory tests the DeleteOneCategory use case method
func (s *CategoryUsecaseTestSuite) TestDeleteOneCategory() {
	// Valid category ID
	s.Run("Valid category ID", func() {
		categoryID := 1

		s.mockRepo.EXPECT().DeleteOneCategory(s.ctx, categoryID).Return(nil)

		err := s.uc.DeleteOneCategory(s.ctx, categoryID)

		assert.NoError(s.T(), err, "Expected no error while deleting category")
	})

	// Invalid category ID
	s.Run("Invalid category ID", func() {
		categoryID := -1

		err := s.uc.DeleteOneCategory(s.ctx, categoryID)

		assert.Error(s.T(), err, "Expected error due to invalid category ID")
	})

	// Category not found
	s.Run("Category not found", func() {
		categoryID := 999

		s.mockRepo.EXPECT().DeleteOneCategory(s.ctx, categoryID).Return(errors.New("category not found"))

		err := s.uc.DeleteOneCategory(s.ctx, categoryID)

		assert.Error(s.T(), err, "Expected error due to category not found")
	})
}

// TestFindListCategory tests the FindListCategory use case method
func (s *CategoryUsecaseTestSuite) TestFindListCategory() {
	params := &request.ListCategoryRequest{
		BaseQuery: request.BaseQuery{
			Page:    1,
			PerPage: 10,
			Keyword: "Test",
		},
	}

	// Valid input
	s.Run("Valid input", func() {
		expectedCategories := []model.Category{
			{Name: "Test Category 1"},
			{Name: "Test Category 2"},
		}

		s.mockRepo.EXPECT().FindListCategory(s.ctx, params).Return(expectedCategories, int64(len(expectedCategories)), nil)

		categories, count, err := s.uc.FindListCategory(s.ctx, params)

		assert.NoError(s.T(), err, "Expected no error while finding list of categories")
		assert.Equal(s.T(), int64(len(expectedCategories)), count, "Expected category count to match")
		assert.Equal(s.T(), expectedCategories, categories, "Expected categories to match")
	})

	// Empty results
	s.Run("Empty results", func() {
		expectedCategories := []model.Category{}

		s.mockRepo.EXPECT().FindListCategory(s.ctx, params).Return(expectedCategories, int64(0), nil)

		categories, count, err := s.uc.FindListCategory(s.ctx, params)

		assert.NoError(s.T(), err, "Expected no error while finding list of categories")
		assert.Equal(s.T(), int64(0), count, "Expected category count to be zero")
		assert.Equal(s.T(), expectedCategories, categories, "Expected categories to be empty")
	})

	// Database error
	s.Run("Database error", func() {
		s.mockRepo.EXPECT().FindListCategory(s.ctx, params).Return(nil, int64(0), errors.New("database error"))

		categories, count, err := s.uc.FindListCategory(s.ctx, params)

		assert.Error(s.T(), err, "Expected error due to database error")
		assert.Equal(s.T(), int64(0), count, "Expected category count to be zero")
		assert.Nil(s.T(), categories, "Expected no categories to be found")
	})
}

// TestFindOneCategory tests the FindOneCategory use case method
func (s *CategoryUsecaseTestSuite) TestFindOneCategory() {
	// Valid category ID
	s.Run("Valid category ID", func() {
		categoryID := 1
		expectedCategory := &model.Category{
			CategoryID: categoryID,
			Name:       "Test Category",
		}

		s.mockRepo.EXPECT().FindOneCategory(s.ctx, categoryID).Return(expectedCategory, nil)

		category, err := s.uc.FindOneCategory(s.ctx, categoryID)

		assert.NoError(s.T(), err, "Expected no error while finding one category")
		assert.NotNil(s.T(), category, "Expected category to be found")
		assert.Equal(s.T(), expectedCategory, category, "Expected category to match")
	})

	// Invalid category ID
	s.Run("Invalid category ID", func() {
		categoryID := -1

		_, err := s.uc.FindOneCategory(s.ctx, categoryID)

		assert.Error(s.T(), err, "Expected error due to invalid category ID")
	})

	// Category not found
	s.Run("Category not found", func() {
		categoryID := 999

		s.mockRepo.EXPECT().FindOneCategory(s.ctx, categoryID).Return(nil, errors.New("category not found"))

		category, err := s.uc.FindOneCategory(s.ctx, categoryID)

		assert.Error(s.T(), err, "Expected error due to category not found")
		assert.Nil(s.T(), category, "Expected no category to be found")
	})
}

// TestCategoryUsecaseSuite runs the Category use case test suite
func TestCategoryUsecaseSuite(t *testing.T) {
	suite.Run(t, new(CategoryUsecaseTestSuite))
}
