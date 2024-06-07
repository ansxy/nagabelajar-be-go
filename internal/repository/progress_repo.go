package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
)

// CreateProgress is a function to create progress
func (r *Repository) CreateProgress(ctx context.Context, data *model.Progress) error {
	return r.BaseRepository.Create(r.db.WithContext(ctx), data)
}

// FindOneProgress is a function to find one progress
func (r *Repository) FindOneProgress(ctx context.Context, query ...interface{}) (*model.Progress, error) {
	var res *model.Progress

	if err := r.BaseRepository.FindOne(r.db.WithContext(ctx).Where(query[0], query[1:]...), &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateProgress is a function to update progress
func (r *Repository) UpdateProgress(ctx context.Context, data *model.Progress) error {
	return r.BaseRepository.Update(r.db.WithContext(ctx), data)
}

// FindListProgress is a function to find list progress
func (r *Repository) FindListProgress(ctx context.Context, courseID int, userID string) ([]model.Progress, int64, error) {
	var progress []model.Progress
	var count int64

	query := r.db.WithContext(ctx).Model(&model.Progress{}).Where("course_id = ? AND user_id = ?", courseID, userID)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Find(&progress).Error; err != nil {
		return nil, 0, err
	}

	return progress, count, nil
}
