package usecase

import (
	"context"
	"strings"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

// CreateCourse implements IFaceUsecase.
func (u *Usecase) CreateCourse(ctx context.Context, data *request.UpsertCourseRequest) error {
	code := GenerateCodeCourse(data.Name)
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

// Generate Course Code
func GenerateCodeCourse(name string) string {
	words := strings.Split(name, " ")
	abbreviation := ""
	for _, word := range words {
		abbreviation += string(word[0])
	}
	return strings.ToUpper(abbreviation)
}
