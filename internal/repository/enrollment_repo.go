package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
)

// CreateEnrollment is a function to create enrollment
func (r *Repository) CreateEnrollment(ctx context.Context, data *model.Enrollment) error {
	return r.BaseRepository.Create(r.db.WithContext(ctx), data)
}
