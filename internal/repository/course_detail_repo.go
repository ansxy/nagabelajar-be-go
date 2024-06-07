package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"gorm.io/gorm"
)

// CreateCourseDetail implements IFaceRepository.
func (r *Repository) CreateCourseDetail(ctx context.Context, data *model.CourseDetail) error {
	return r.BaseRepository.Create(r.db.WithContext(ctx), data)
}

// FindCourseDetail implements IFaceRepository.
func (r *Repository) FindOneCourseDetail(ctx context.Context, courseDetailID int) (*model.CourseDetail, error) {
	var data model.CourseDetail

	if err := r.BaseRepository.FindOne(
		r.db.WithContext(ctx).Where("course_detail_id = ?", courseDetailID).Preload("Content", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
			Preload("Content.SubContent").Preload("Content.Exercise").Preload("Content.Practice").Preload("Assigment"),
		&data); err != nil {
		return nil, err
	}

	return &data, nil
}
