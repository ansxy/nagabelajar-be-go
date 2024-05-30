package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
)

// CreateMedia implements IFaceRepository.
func (repo *Repository) CreateMedia(ctx context.Context, data *model.Media) error {
	return repo.BaseRepository.Create(repo.db.WithContext(ctx), data)
}
