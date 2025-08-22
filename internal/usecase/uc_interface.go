package usecase

import (
	"context"
	"mime/multipart"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth/artifact"
)

type IFaceUsecase interface {
	//User
	RegisterUser(ctx context.Context, data *request.UpsertUserRequest) error
	Login(ctx context.Context, data *request.LoginRequest) (*model.User, error)
	LoginWithGoogle(ctx context.Context, data *request.LoginWithGoogleRequest) (*model.User, error)
	FindOneUser(ctx context.Context, data *request.LoginRequest) (*model.User, error)

	//Category
	CreateCategory(ctx context.Context, data *request.UpsertCategoryRequest) error
	FindListCategory(ctx context.Context, params *request.ListCategoryRequest) ([]model.Category, int64, error)
	DeleteOneCategory(ctx context.Context, categoryID int) error
	FindOneCategory(ctx context.Context, categoryID int) (*model.Category, error)

	//Course
	CreateCourse(ctx context.Context, data *request.UpsertCourseRequest) error
	DeleteCourse(ctx context.Context, courseID int) error
	FindListCourse(ctx context.Context, params *request.ListCourseRequest) ([]model.Course, int64, error)
	FindOneCourse(ctx context.Context, params *request.GetOneCourseRequest) (*model.Course, error)
	EnrollCourse(ctx context.Context, data *request.EnrollCourseRequest) error

	//CourseDetail
	CreateCourseDetail(ctx context.Context, data *request.UpserCourseDetailRequest) error
	FindCourseDetail(ctx context.Context, courseID int) (*model.CourseDetail, error)

	//Transaction
	CreateTransaction(ctx context.Context, data *request.InsertTransaction) error

	//File
	UploadFile(ctx context.Context, header *multipart.FileHeader) error

	//Certificate
	ValidateCertificate(ctx context.Context, file *multipart.FileHeader) error
	ValidateCertificateByAddress(ctx context.Context, address string) (*goeth.CertificateOfCompletionCertificateData, error)
	CreateCertificate(ctx context.Context, req *request.CreateCertificateRequest) error
	GetListCertificate(ctx context.Context, params *request.ListCertificateRequest) ([]model.Certificate, int64, error)

	//Media
	UploadMedia(ctx context.Context, data *multipart.FileHeader) (*model.Media, error)

	//Progress
	UpdateProgress(ctx context.Context, data *request.UpdateProgressRequest) error
}
