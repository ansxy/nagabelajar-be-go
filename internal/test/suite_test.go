package test

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mock_repo "github.com/ansxy/nagabelajar-be-go/internal/mock"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UsecaseTestSuite struct {
	suite.Suite
	uc         usecase.IFaceUsecase
	mockRepo   *mock_repo.MockIFaceRepository
	ctx        context.Context
	goMockCtrl *gomock.Controller
	dbMock     sqlmock.Sqlmock
}

func (s *UsecaseTestSuite) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		log.Println("Error creating mock database", err.Error())
	}

	dbClient, err := gorm.Open(postgres.New(postgres.Config{Conn: dbMock}))
	require.NoError(s.T(), err)
	s.dbMock = mock

	s.ctx = context.TODO()
	s.goMockCtrl = gomock.NewController(s.T())
	s.mockRepo = mock_repo.NewMockIFaceRepository(s.goMockCtrl)

	s.uc = usecase.NewUsecase(&usecase.Usecase{
		Repo: s.mockRepo,
		DB:   dbClient,
	})
}

func (s *UsecaseTestSuite) TearDownTest() {
	s.goMockCtrl.Finish()
}

func (s *UsecaseTestSuite) TestCreateUser() {
	// Setup expectations for the mock repository
	s.mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
	err := s.uc.RegisterUser(s.ctx, &request.UpsertUserRequest{})

	// Verify the result
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), err)

}

// func (s *UsecaseTestSuite) TestLogin() {
// 	// Setup expectations for the mock repository
// 	expectedUser := &model.User{Email: "ansarfadillah212@gmail.com", Name: "ansar"}
// 	s.mockRepo.EXPECT().FindOneUser(gomock.Any(), gomock.Any()).Return(expectedUser, nil)

// 	// Call the Login method of the use case
// 	user, err := s.uc.Login(s.ctx, &request.LoginRequest{Email: "ansarfadillah212@gmail.com", Password: "ansar123"})

// 	// Verify the result
// 	assert.NoError(s.T(), err)
// 	assert.NotNil(s.T(), user)
// 	assert.Equal(s.T(), expectedUser, user)
// }

func (s *UsecaseTestSuite) TestRegister() {
	// Setup expectations for the mock repository
	s.mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
	err := s.uc.RegisterUser(s.ctx, &request.UpsertUserRequest{})

	// Verify the result
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), err)
}

// Add more test methods for other use case methods...

func TestUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UsecaseTestSuite))
}
