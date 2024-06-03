package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"gorm.io/gorm"
)

// CreateCourse implements IFaceRepository.
func (repo *Repository) CreateCourse(ctx context.Context, data *model.Course) error {
	return repo.BaseRepository.Create(repo.db.WithContext(ctx), data)
}

// DeleteOneCourse implements IFaceRepository.
func (repo *Repository) DeleteOneCourse(ctx context.Context, courseID int) error {
	return repo.BaseRepository.Delete(repo.db.WithContext(ctx), &model.Course{
		CourseID: courseID,
	})
}

// FindListCourse implements IFaceRepository.
func (repo *Repository) FindListCourse(ctx context.Context, params *request.ListCourseRequest) ([]model.Course, int64, error) {
	var course []model.Course
	var count int64

	query := repo.db.WithContext(ctx).Model(&model.Course{}).Preload("Category").Preload("Media")
	if params.Keyword != "" {
		lowerCaseKeyword := strings.ToLower(params.Keyword)
		query = query.Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", lowerCaseKeyword))
	}

	if params.Sort != "" {
		query = query.Order(params.Sort)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(params.PerPage).Offset((params.Page - 1) * params.PerPage).Find(&course).Error; err != nil {
		return nil, 0, err
	}

	return course, count, nil
}

// FindOneCourse implements IFaceRepository.
func (repo *Repository) FindOneCourse(ctx context.Context, params *request.GetOneCourseRequest) (*model.Course, error) {
	var res model.Course
	isEnrolled := false
	query := repo.db.WithContext(ctx).Model(&model.Course{}).Where("tr_courses.course_id = ?", params.CourseID).Preload("Category").Preload("Media").Preload("CourseDetail").Preload("Progress")

	if params.UserID != nil {
		query = query.Joins("LEFT JOIN tr_enrollments ON tr_enrollments.course_id = tr_courses.course_id AND tr_enrollments.user_id = ?", *params.UserID).
			Preload("Enrollment", func(db *gorm.DB) *gorm.DB {
				return db.Unscoped().Where("user_id = ?", *params.UserID)
			})
	}

	if err := query.First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.IsEnrolled = &isEnrolled
			return &res, nil
		}
		return nil, err
	}

	if res.Enrollment == nil {
		res.IsEnrolled = &isEnrolled
	}

	isEnrolled = true
	res.IsEnrolled = &isEnrolled

	return &res, nil
}

// UpdateCourse implements IFaceRepository.
func (repo *Repository) UpdateCourse(ctx context.Context, data *model.Course) error {
	panic("implement me")
}
