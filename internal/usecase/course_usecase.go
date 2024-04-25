package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	custom_string "github.com/ansxy/nagabelajar-be-go/pkg/string"
)

// CreateCourse implements IFaceUsecase.
func (u *Usecase) CreateCourse(ctx context.Context, data *request.UpsertCourseRequest) error {
	var count int64
	count, err := u.Repo.CountCategoryCourse(ctx, data.CategoryID)

	if err != nil {
		return err
	}

	code := custom_string.GenerateCodeCourse(data.Name, int(count))
	course := &model.Course{
		Name:        data.Name,
		Code:        code,
		CategoryID:  data.CategoryID,
		Price:       data.Price,
		Description: data.Description,
		IsPaid:      data.IsPaid,
		IsArchived:  data.IsArchived,
	}

	return u.Repo.CreateCourse(ctx, course)
}

// FindListCourse implements IFaceUsecase
func (u *Usecase) FindListCourse(ctx context.Context, params *request.ListCourseRequest) ([]model.Course, int64, error) {
	return u.Repo.FindListCourse(ctx, params)
}

// DeleteCourse implements IFaceUsecase
func (u *Usecase) DeleteCourse(ctx context.Context, courseID int) error {
	return u.Repo.DeleteOneCourse(ctx, courseID)
}

// FindOneCourse implements IFaceUsecase.
func (u *Usecase) FindOneCourse(ctx context.Context, courseID string) (*model.Course, error) {
	return u.Repo.FindOneCourse(ctx, courseID)
}
