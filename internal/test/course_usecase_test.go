package test

import (
	"errors"
	"testing"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

// CourseUsecaseTestSuite is a test suite for the Course use case
type CourseUsecaseTestSuite struct {
	BaseTestSuite
	uc usecase.IFaceUsecase
}

// SetupTest sets up the test environment for Course use case
func (s *CourseUsecaseTestSuite) SetupTest() {
	s.BaseTestSuite.SetupTest()
	s.uc = usecase.NewUsecase(&usecase.Usecase{
		Repo:      s.mockRepo,
		DB:        nil, // assuming DB is not needed for the usecase or provide actual db connection if needed
		Validator: s.validate,
	})
}

// TestCreateCourse tests the CreateCourse use case method
func (s *CourseUsecaseTestSuite) TestCreateCourse() {
	// Valid input
	s.Run("Valid input", func() {
		req := &request.UpsertCourseRequest{
			Name:        "Test Course",
			CategoryID:  1,
			Price:       100,
			Description: "Test Description",
			IsPaid:      func(b bool) *bool { return &b }(true),
			IsArchived:  func(b bool) *bool { return &b }(false),
			Author:      "Test Author",
			MediaID:     uuid.NewString(),
		}

		s.mockRepo.EXPECT().CountCategoryCourse(s.ctx, req.CategoryID).Return(int64(0), nil)
		s.mockRepo.EXPECT().CreateCourse(s.ctx, gomock.Any()).Return(nil)

		err := s.uc.CreateCourse(s.ctx, req)

		assert.NoError(s.T(), err, "Expected no error while creating course")
	})

	// Missing required fields
	s.Run("Missing required fields", func() {
		req := &request.UpsertCourseRequest{
			Name: "Test Course",
			// CategoryID:  0, // Missing CategoryID
			Price:       100,
			Description: "Test Description",
			IsPaid:      func(b bool) *bool { return &b }(true),
			IsArchived:  func(b bool) *bool { return &b }(false),
			Author:      "Test Author",
			MediaID:     uuid.NewString(),
		}

		// Validate struct
		err := s.validateStruct(req)
		assert.Error(s.T(), err, "Expected error due to missing required fields")
	})

	// Database error
	s.Run("Database error", func() {
		req := &request.UpsertCourseRequest{
			Name:        "Test Course",
			CategoryID:  1,
			Price:       100,
			Description: "Test Description",
			IsPaid:      func(b bool) *bool { return &b }(true),
			IsArchived:  func(b bool) *bool { return &b }(false),
			Author:      "Test Author",
			MediaID:     uuid.NewString(),
		}

		// Validate struct
		err := s.validateStruct(req)
		require.NoError(s.T(), err)

		s.mockRepo.EXPECT().CountCategoryCourse(s.ctx, req.CategoryID).Return(int64(0), nil)
		s.mockRepo.EXPECT().CreateCourse(s.ctx, gomock.Any()).Return(errors.New("database error"))

		err = s.uc.CreateCourse(s.ctx, req)

		assert.Error(s.T(), err, "Expected error due to database error")
	})
}

// TestFindOneCourse tests the FindOneCourse use case method
func (s *CourseUsecaseTestSuite) TestFindOneCourse() {
	// Valid course ID
	s.Run("Valid course ID", func() {
		courseID := "1"
		expectedCourse := &model.Course{
			CourseID: 1,
			Name:     "Test Course",
		}

		req := &request.GetOneCourseRequest{
			CourseID: courseID,
		}

		s.mockRepo.EXPECT().FindOneCourse(s.ctx, req).Return(expectedCourse, nil)

		course, err := s.uc.FindOneCourse(s.ctx, req)

		assert.NoError(s.T(), err, "Expected no error while finding one course")
		assert.NotNil(s.T(), course, "Expected course to be found")
		assert.Equal(s.T(), expectedCourse, course, "Expected course to match")
	})

	// Invalid course ID
	s.Run("Invalid course ID", func() {
		courseID := "invalid"

		req := &request.GetOneCourseRequest{
			CourseID: courseID,
		}

		s.mockRepo.EXPECT().FindOneCourse(s.ctx, req).Return(nil, errors.New("invalid course ID"))

		course, err := s.uc.FindOneCourse(s.ctx, req)

		assert.Error(s.T(), err, "Expected error due to invalid course ID")
		assert.Nil(s.T(), course, "Expected no course to be found")
	})

	// Course not found
	s.Run("Course not found", func() {
		courseID := "999"

		req := &request.GetOneCourseRequest{
			CourseID: courseID,
		}

		s.mockRepo.EXPECT().FindOneCourse(s.ctx, req).Return(nil, gorm.ErrRecordNotFound)

		course, err := s.uc.FindOneCourse(s.ctx, req)

		assert.Error(s.T(), err, "Expected error due to course not found")
		assert.Nil(s.T(), course, "Expected no course to be found")
	})
}

// TestFindListCourse tests the FindListCourse use case method
func (s *CourseUsecaseTestSuite) TestFindListCourse() {
	params := &request.ListCourseRequest{
		Name: "test",
		BaseQuery: request.BaseQuery{
			Page:    1,
			PerPage: 10,
		},
	}

	// Valid input
	s.Run("Valid input", func() {
		expectedCourses := []model.Course{
			{Name: "Test Course 1"},
			{Name: "Test Course 2"},
		}

		s.mockRepo.EXPECT().FindListCourse(s.ctx, params).Return(expectedCourses, int64(len(expectedCourses)), nil)

		courses, count, err := s.uc.FindListCourse(s.ctx, params)

		assert.NoError(s.T(), err, "Expected no error while finding list of courses")
		assert.Equal(s.T(), int64(len(expectedCourses)), count, "Expected course count to match")
		assert.Equal(s.T(), expectedCourses, courses, "Expected courses to match")
	})

	// Empty results
	s.Run("Empty results", func() {
		expectedCourses := []model.Course{}

		s.mockRepo.EXPECT().FindListCourse(s.ctx, params).Return(expectedCourses, int64(0), nil)

		courses, count, err := s.uc.FindListCourse(s.ctx, params)

		assert.NoError(s.T(), err, "Expected no error while finding list of courses")
		assert.Equal(s.T(), int64(0), count, "Expected course count to be zero")
		assert.Equal(s.T(), expectedCourses, courses, "Expected courses to be empty")
	})

	// Database error
	s.Run("Database error", func() {
		s.mockRepo.EXPECT().FindListCourse(s.ctx, params).Return(nil, int64(0), errors.New("database error"))

		courses, count, err := s.uc.FindListCourse(s.ctx, params)

		assert.Error(s.T(), err, "Expected error due to database error")
		assert.Equal(s.T(), int64(0), count, "Expected course count to be zero")
		assert.Nil(s.T(), courses, "Expected no courses to be found")
	})
}

// TestDeleteCourse tests the DeleteCourse use case method
func (s *CourseUsecaseTestSuite) TestDeleteCourse() {
	// Valid course ID
	s.Run("Valid course ID", func() {
		courseID := 1

		s.mockRepo.EXPECT().DeleteOneCourse(s.ctx, courseID).Return(nil)

		err := s.uc.DeleteCourse(s.ctx, courseID)

		assert.NoError(s.T(), err, "Expected no error while deleting course")
	})

	// Invalid course ID
	s.Run("Invalid course ID", func() {
		courseID := -1

		s.mockRepo.EXPECT().DeleteOneCourse(s.ctx, courseID).Return(errors.New("invalid course ID"))

		err := s.uc.DeleteCourse(s.ctx, courseID)

		assert.Error(s.T(), err, "Expected error due to invalid course ID")
	})

	// Course not found
	s.Run("Course not found", func() {
		courseID := 999

		s.mockRepo.EXPECT().DeleteOneCourse(s.ctx, courseID).Return(gorm.ErrRecordNotFound)

		err := s.uc.DeleteCourse(s.ctx, courseID)

		assert.Error(s.T(), err, "Expected error due to course not found")
	})
}

// TestCourseUsecaseSuite runs the Course use case test suite
func TestCourseUsecaseSuite(t *testing.T) {
	suite.Run(t, new(CourseUsecaseTestSuite))
}
