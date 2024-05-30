package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

//go:generate mockgen -destination=../mock/repository_mock.go -package=mock_repo -source=repo_interface.go
type IFaceRepository interface {
	//User
	CreateUser(ctx context.Context, data *model.User) error
	FindOneUser(ctx context.Context, query ...interface{}) (*model.User, error)
	UpdateUser(ctx context.Context, data *model.User) error
	FindUsers(ctx context.Context, params *request.ListUserRequest) ([]model.User, int64, error)
	//CategoryP
	CreateCategory(ctx context.Context, data *model.Category) error
	FindListCategory(ctx context.Context, params *request.ListCategoryRequest) ([]model.Category, int64, error)
	DeleteOneCategory(ctx context.Context, categoryID int) error
	FindOneCategory(ctx context.Context, categoryID int) (*model.Category, error)
	CountCategoryCourse(ctx context.Context, categoryID int) (int64, error)

	//Course
	CreateCourse(ctx context.Context, data *model.Course) error
	FindListCourse(ctx context.Context, params *request.ListCourseRequest) ([]model.Course, int64, error)
	FindOneCourse(ctx context.Context, courseID string) (*model.Course, error)
	DeleteOneCourse(ctx context.Context, courseID int) error
	UpdateCourse(ctx context.Context, data *model.Course) error

	//DetailCourse
	CreateCourseDetail(ctx context.Context, data *model.CourseDetail) error
	FindOneCourseDetail(ctx context.Context, course_detail_id int) (*model.CourseDetail, error)

	//Transaction
	CreateTransaction(ctx context.Context, data *model.Transaction) error

	//Certificate
	CreateCertificate(ctx context.Context, data *model.Certificate) error
	FindOneCertificate(ctx context.Context, query ...interface{}) (*model.Certificate, error)

	//Media
	CreateMedia(ctx context.Context, data *model.Media) error
}
