package usecase

import (
	"context"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	custom_string "github.com/ansxy/nagabelajar-be-go/pkg/string"
	"github.com/google/uuid"
)

// CreateCourse implements IFaceUsecase.
func (u *Usecase) CreateCourse(ctx context.Context, data *request.UpsertCourseRequest) error {
	var count int64
	count, err := u.Repo.CountCategoryCourse(ctx, data.CategoryID)

	if err != nil {
		return err
	}

	code := custom_string.GenerateCodeCourse(data.Name, int(count))
	mediaID, _ := uuid.Parse(data.MediaID)
	course := &model.Course{
		Name:        data.Name,
		Code:        code,
		CategoryID:  data.CategoryID,
		Price:       data.Price,
		Description: data.Description,
		IsPaid:      data.IsPaid,
		IsArchived:  data.IsArchived,
		Author:      data.Author,
		MediaID:     mediaID,
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
func (u *Usecase) FindOneCourse(ctx context.Context, params *request.GetOneCourseRequest) (*model.Course, error) {
	return u.Repo.FindOneCourse(ctx, params)
}

// EnrollCourse implements IFaceUsecase.
func (u *Usecase) EnrollCourse(ctx context.Context, data *request.EnrollCourseRequest) error {

	var course *model.Course

	course, err := u.Repo.FindOneCourse(ctx, &request.GetOneCourseRequest{
		CourseID: data.CourseID,
	})

	if err != nil {
		return err
	}

	for _, content := range course.CourseDetail {
		courseID, err := strconv.Atoi(data.CourseID)
		if err != nil {
			return err
		}
		progress := &model.Progress{
			UserID:         data.UserID.String(),
			CourseID:       courseID,
			CourseDetailID: content.CourseDetailID,
			IsFinished:     false,
		}

		err = u.Repo.CreateProgress(ctx, progress)
		if err != nil {
			return err
		}
	}

	enrollment := &model.Enrollment{
		UserID:   data.UserID.String(),
		CourseID: data.CourseID,
	}

	return u.Repo.CreateEnrollment(ctx, enrollment)
}
