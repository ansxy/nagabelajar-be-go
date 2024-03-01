package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

type IFaceUsecase interface {
	//User
	RegisterUser(ctx context.Context, data *request.UpsertUserRequest) error
	Login(ctx context.Context, data *request.LoginRequest) (*model.User, error)

	//Category
	CreateCategory(ctx context.Context, data *request.UpsertCategoryRequest) error
	FindListCategory(ctx context.Context, params *request.ListCategoryRequest) ([]model.Category, int64, error)
	DeleteOneCategory(ctx context.Context, categoryID int) error
	FindOneCategory(ctx context.Context, categoryID int) (*model.Category, error)

	//Course
	CreateCourse(ctx context.Context, data *request.UpsertCourseRequest) error
	DeleteCourse(ctx context.Context, courseID int) error
	FindListCourse(ctx context.Context, params *request.ListCourseRequest) ([]model.Course, int64, error)
	FindOneCourse(ctx context.Context, courseID int) (*model.Course, error)

	//Transaction
	CreateTransaction(ctx context.Context, data *request.InsertTransaction) error
}
