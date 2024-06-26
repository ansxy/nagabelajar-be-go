package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

//go:generate mockgen -source=base_repo.go -destination=mocks/mock_base_repo.go -package=mocks
type IFaceRepository interface {
	//User
	CreateUser(ctx context.Context, data *model.User) error
	FindOneUser(ctx context.Context, query ...interface{}) (*model.User, error)
	UpdateUser(ctx context.Context, data *model.User) error

	//Category
	CreateCategory(ctx context.Context, data *model.Category) error
	FindListCategory(ctx context.Context, params *request.ListCategoryRequest) ([]model.Category, int64, error)
	DeleteOneCategory(ctx context.Context, categoryID int) error
	FindOneCategory(ctx context.Context, categoryID int) (*model.Category, error)

	//Course
	CreateCourse(ctx context.Context, data *model.Course) error
	FindListCourse(ctx context.Context, params *request.ListCourseRequest) ([]model.Course, int64, error)
	FindOneCourse(ctx context.Context, courseID int) (*model.Course, error)
	DeleteOneCourse(ctx context.Context, courseID int) error
	UpdateCourse(ctx context.Context, data *model.Course) error

	//Transaction
	CreateTransaction(ctx context.Context, data *model.Transaction) error
}
