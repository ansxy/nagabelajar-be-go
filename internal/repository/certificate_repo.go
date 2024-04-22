package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
)

// CreateCertificate implements IFaceRepository.
func (repo *Repository) CreateCertificate(ctx context.Context, data *model.Certificate) error {
	return repo.BaseRepository.Create(repo.db.WithContext(ctx), data)
}

// FindOneCertificate implements IFaceRepository.
func (repo *Repository) FindOneCertificate(ctx context.Context, query ...interface{}) (*model.Certificate, error) {
	var res *model.Certificate

	if err := repo.BaseRepository.FindOne(repo.db.WithContext(ctx).Where(query[0], query[:1]...), &res); err != nil {
		return nil, err
	}

	return res, nil
}
