package test

import (
	"context"
	"errors"
	"testing"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/ansxy/nagabelajar-be-go/pkg/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

// UserUsecaseTestSuite is a test suite for the User use case
type UserUsecaseTestSuite struct {
	BaseTestSuite
	uc usecase.IFaceUsecase
}

// SetupTest sets up the test environment for User use case
func (s *UserUsecaseTestSuite) SetupTest() {
	s.BaseTestSuite.SetupTest()
	s.uc = usecase.NewUsecase(&usecase.Usecase{
		Repo: s.mockRepo,
	})
}

// MockHashPassword is a helper function to simulate hashing errors
func MockHashPassword(success bool) func(string) (string, error) {
	return func(password string) (string, error) {
		if success {
			return hash.HashPassword(password)
		}
		return "", errors.New("hashing error")
	}
}

// TestRegisterUser tests the RegisterUser use case method
func (s *UserUsecaseTestSuite) TestRegisterUser() {
	s.Run("Register user", func() {
		req := &request.UpsertUserRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
		}

		hashedPassword, err := MockHashPassword(true)(req.Password)
		require.NoError(s.T(), err)

		user := &model.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: hashedPassword,
		}

		s.mockRepo.EXPECT().CreateUser(s.ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, userArg *model.User) error {
			assert.Equal(s.T(), user.Name, userArg.Name)
			assert.Equal(s.T(), user.Email, userArg.Email)
			assert.NoError(s.T(), hash.ComparePassword(userArg.Password, req.Password), "Expected password to be hashed correctly")
			return nil
		})

		err = s.uc.RegisterUser(s.ctx, req)
		assert.NoError(s.T(), err, "Expected no error while registering user")
	})

	s.Run("Register user with hashing error", func() {
		req := &request.UpsertUserRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
		}

		hashPasswordFunc := MockHashPassword(false)
		_, err := hashPasswordFunc(req.Password)
		assert.Error(s.T(), err, "Expected hashing error")
	})
}

// TestLogin tests the Login use case method
func (s *UserUsecaseTestSuite) TestLogin() {
	s.Run("Login user", func() {
		req := &request.LoginRequest{
			Email:    "test@example.com",
			Password: "password",
		}

		hashedPassword, err := MockHashPassword(true)(req.Password)
		require.NoError(s.T(), err)

		user := &model.User{
			Email:    req.Email,
			Password: hashedPassword,
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(user, nil)

		loggedInUser, err := s.uc.Login(s.ctx, req)
		assert.NoError(s.T(), err, "Expected no error while logging in")
		assert.Equal(s.T(), user, loggedInUser, "Expected user to match")
	})

	s.Run("Login user with incorrect password", func() {
		req := &request.LoginRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}

		hashedPassword, err := MockHashPassword(true)("password")
		require.NoError(s.T(), err)

		user := &model.User{
			Email:    req.Email,
			Password: hashedPassword,
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(user, nil)

		loggedInUser, err := s.uc.Login(s.ctx, req)
		assert.Error(s.T(), err, "Expected error due to incorrect password")
		assert.Nil(s.T(), loggedInUser, "Expected no user to be logged in")
	})
}

// TestLoginWithGoogle tests the LoginWithGoogle use case method
func (s *UserUsecaseTestSuite) TestLoginWithGoogle() {
	s.Run("Login with Google existing user", func() {
		req := &request.LoginWithGoogleRequest{
			FirebaseID: "test-firebase-id",
			Email:      "test@example.com",
		}

		user := &model.User{
			FirebaseID: req.FirebaseID,
			Email:      req.Email,
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(user, nil)

		loggedInUser, err := s.uc.LoginWithGoogle(s.ctx, req)
		assert.NoError(s.T(), err, "Expected no error while logging in with Google")
		assert.Equal(s.T(), user, loggedInUser, "Expected user to match")
	})

	s.Run("Login with Google new user", func() {
		req := &request.LoginWithGoogleRequest{
			FirebaseID: "test-firebase-id",
			Email:      "test@example.com",
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
		s.mockRepo.EXPECT().CreateUser(s.ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, userArg *model.User) error {
			assert.Equal(s.T(), req.FirebaseID, userArg.FirebaseID)
			assert.Equal(s.T(), req.Email, userArg.Email)
			assert.True(s.T(), userArg.IsGoogle)
			assert.Equal(s.T(), constant.Role.User, userArg.Role)
			return nil
		})

		loggedInUser, err := s.uc.LoginWithGoogle(s.ctx, req)
		assert.NoError(s.T(), err, "Expected no error while logging in with Google")
		assert.Equal(s.T(), req.FirebaseID, loggedInUser.FirebaseID, "Expected Firebase ID to match")
		assert.Equal(s.T(), req.Email, loggedInUser.Email, "Expected email to match")
	})
}

// TestFindOneUser tests the FindOneUser use case method
func (s *UserUsecaseTestSuite) TestFindOneUser() {
	s.Run("Find one user", func() {
		req := &request.LoginRequest{
			FirebaseID: "test-firebase-id",
		}

		user := &model.User{
			FirebaseID: req.FirebaseID,
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, "firebase_id = ?", req.FirebaseID).Return(user, nil)

		foundUser, err := s.uc.FindOneUser(s.ctx, req)
		assert.NoError(s.T(), err, "Expected no error while finding one user")
		assert.Equal(s.T(), user, foundUser, "Expected user to match")
	})

	s.Run("Find one user not found", func() {
		req := &request.LoginRequest{
			FirebaseID: "test-firebase-id",
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, "firebase_id = ?", req.FirebaseID).Return(nil, gorm.ErrRecordNotFound)

		foundUser, err := s.uc.FindOneUser(s.ctx, req)
		assert.Error(s.T(), err, "Expected error due to user not found")
		assert.Nil(s.T(), foundUser, "Expected no user to be found")
	})
}

// TestUserUsecaseSuite runs the User use case test suite
func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
