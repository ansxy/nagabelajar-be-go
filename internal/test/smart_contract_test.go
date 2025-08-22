package test

import (
	"errors"
	"testing"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

// CertificateUsecaseTestSuite is a test suite for the Certificate use case
type CertificateUsecaseTestSuite struct {
	BaseTestSuite
	uc usecase.IFaceUsecase
}

// SetupTest sets up the test environment for Certificate use case
func (s *CertificateUsecaseTestSuite) SetupTest() {
	s.BaseTestSuite.SetupTest()
	s.uc = usecase.NewUsecase(&usecase.Usecase{
		Repo:    s.mockRepo,
		Service: s.serviceMock,
	})
}

// TestCreateCertificate tests the CreateCertificate use case method
func (s *CertificateUsecaseTestSuite) TestCreateCertificate() {
	s.Run("Create certificate", func() {
		// req := &request.CreateCertificateRequest{
		// 	FirebaseID: "test-firebase-id",
		// 	CourseId:   1,
		// }

		user := &model.User{
			FirebaseID: "test-firebase-id",
			Name:       "Test User",
			Email:      "test@example.com",
		}

		course := &model.Course{
			CourseID: 1,
			Name:     "Test Course",
			Code:     "TC01",
		}

		newCertAddress := common.HexToAddress("0x123")
		tx := &types.Transaction{}

		// Mock expectations
		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(user, nil)
		s.mockRepo.EXPECT().FindOneCourse(s.ctx, gomock.Any()).Return(course, nil)
		s.serviceMock.EXPECT().CreateCertificate(s.ctx, gomock.Any(), user.Name, user.Email, course.Name, course.Code).Return(&newCertAddress, tx, nil)
		s.serviceMock.EXPECT().QRCodeGenerator(gomock.Any()).Return("qr-code", nil)
		s.serviceMock.EXPECT().CertificatePDF(s.ctx, tx, gomock.Any(), newCertAddress).Return([]byte("pdf-content"), nil)
		s.serviceMock.EXPECT().Md5Reader(s.ctx, gomock.Any(), gomock.Any()).Return("md5-hash", nil)
		s.serviceMock.EXPECT().UpdateMd5Certificate(s.ctx, gomock.Any(), &newCertAddress, "md5-hash").Return(nil)
		s.mockRepo.EXPECT().CreateCertificate(s.ctx, gomock.Any()).Return(nil)

		// err := s.uc.CreateCertificate(s.ctx, req)

		// assert.NoError(s.T(), err, "Expected no error while creating certificate")
	})

	// Add more sub-tests for different scenarios such as errors, invalid inputs, etc.
	s.Run("Error finding user", func() {
		req := &request.CreateCertificateRequest{
			FirebaseID: "test-firebase-id",
			CourseId:   1,
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(nil, errors.New("user not found"))

		err := s.uc.CreateCertificate(s.ctx, req)

		assert.Error(s.T(), err, "Expected error due to user not found")
	})

	s.Run("Error finding course", func() {
		req := &request.CreateCertificateRequest{
			FirebaseID: "test-firebase-id",
			CourseId:   1,
		}

		user := &model.User{
			FirebaseID: "test-firebase-id",
			Name:       "Test User",
			Email:      "test@example.com",
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(user, nil)
		s.mockRepo.EXPECT().FindOneCourse(s.ctx, gomock.Any()).Return(nil, errors.New("course not found"))

		err := s.uc.CreateCertificate(s.ctx, req)

		assert.Error(s.T(), err, "Expected error due to course not found")
	})

	s.Run("Error creating certificate on blockchain", func() {
		req := &request.CreateCertificateRequest{
			FirebaseID: "test-firebase-id",
			CourseId:   1,
		}

		user := &model.User{
			FirebaseID: "test-firebase-id",
			Name:       "Test User",
			Email:      "test@example.com",
		}

		course := &model.Course{
			CourseID: 1,
			Name:     "Test Course",
			Code:     "TC01",
		}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(user, nil)
		s.mockRepo.EXPECT().FindOneCourse(s.ctx, gomock.Any()).Return(course, nil)
		s.serviceMock.EXPECT().CreateCertificate(s.ctx, gomock.Any(), user.Name, user.Email, course.Name, course.Code).Return(nil, nil, errors.New("blockchain error"))

		err := s.uc.CreateCertificate(s.ctx, req)

		assert.Error(s.T(), err, "Expected error due to blockchain error")
	})

	s.Run("Error generating QR code", func() {
		req := &request.CreateCertificateRequest{
			FirebaseID: "test-firebase-id",
			CourseId:   1,
		}

		user := &model.User{
			FirebaseID: "test-firebase-id",
			Name:       "Test User",
			Email:      "test@example.com",
		}

		course := &model.Course{
			CourseID: 1,
			Name:     "Test Course",
			Code:     "TC01",
		}

		newCertAddress := common.HexToAddress("0x123")
		tx := &types.Transaction{}

		s.mockRepo.EXPECT().FindOneUser(s.ctx, gomock.Any()).Return(user, nil)
		s.mockRepo.EXPECT().FindOneCourse(s.ctx, gomock.Any()).Return(course, nil)
		s.serviceMock.EXPECT().CreateCertificate(s.ctx, gomock.Any(), user.Name, user.Email, course.Name, course.Code).Return(&newCertAddress, tx, nil)
		s.serviceMock.EXPECT().QRCodeGenerator(gomock.Any()).Return(nil, errors.New("qr code error"))

		err := s.uc.CreateCertificate(s.ctx, req)

		assert.Error(s.T(), err, "Expected error due to QR code generation error")
	})
}

// TestCertificateUsecaseSuite runs the Certificate use case test suite
func TestCertificateUsecaseSuite(t *testing.T) {
	suite.Run(t, new(CertificateUsecaseTestSuite))
}
