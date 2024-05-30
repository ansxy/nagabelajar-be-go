package test

import (
	"context"
	"errors"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	mock_repo "github.com/ansxy/nagabelajar-be-go/internal/mock"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BaseTestSuite is a base test suite for setting up common test environment
type BaseTestSuite struct {
	suite.Suite
	mockRepo   *mock_repo.MockIFaceRepository
	ctx        context.Context
	goMockCtrl *gomock.Controller
	dbMock     sqlmock.Sqlmock
	validate   *validator.Validate
}

// SetupTest sets up the common test environment
func (s *BaseTestSuite) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	require.NoError(s.T(), err, "Failed to create mock database")

	_, err = gorm.Open(postgres.New(postgres.Config{Conn: dbMock}))
	require.NoError(s.T(), err, "Failed to open GORM DB")

	s.dbMock = mock
	s.ctx = context.TODO()
	s.goMockCtrl = gomock.NewController(s.T())
	s.mockRepo = mock_repo.NewMockIFaceRepository(s.goMockCtrl)
	s.validate = validator.New()
	s.validate.RegisterValidation("positive", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() > 0
	})
}

// TearDownTest tears down the common test environment
func (s *BaseTestSuite) TearDownTest() {
	s.goMockCtrl.Finish()
}

// validateStruct is a helper function to validate struct using go-playground/validator
func (s *BaseTestSuite) validateStruct(req interface{}) error {
	err := s.validate.Struct(req)
	if err != nil {
		var missingFields []string
		for _, err := range err.(validator.ValidationErrors) {
			missingFields = append(missingFields, err.Field())
		}
		return errors.New("fields " + strings.Join(missingFields, ", ") + " are required")
	}
	return nil
}
