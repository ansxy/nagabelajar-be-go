package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
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

	query := repo.db.WithContext(ctx).Model(&model.Course{})

	if params.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+params.Keyword+"%")
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
func (repo *Repository) FindOneCourse(ctx context.Context, courseID int) (*model.Course, error) {
	var res *model.Course

	if err := repo.BaseRepository.FindOne(repo.db.WithContext(ctx).Where("course_id = ?", courseID), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateCourse implements IFaceRepository.
func (repo *Repository) UpdateCourse(ctx context.Context, data *model.Course) error {
	panic("")
}
