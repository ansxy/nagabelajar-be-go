package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
)

// CreateTransaction implements IFaceRepository.
func (repo *Repository) CreateTransaction(ctx context.Context, data *model.Transaction) error {
	return repo.BaseRepository.Create(repo.db.WithContext(ctx), data)
}
