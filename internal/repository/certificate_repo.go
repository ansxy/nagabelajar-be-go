package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
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

// FindListCertificate implements IFaceRepository.
func (repo *Repository) FindListCertificate(ctx context.Context, params *request.ListCertificateRequest) ([]model.Certificate, int64, error) {
	var certificate []model.Certificate
	var count int64

	query := repo.db.WithContext(ctx).Model(&model.Certificate{}).Preload("Course")

	if params.UserID != nil {
		query = query.Where("user_id = ?", *params.UserID)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(params.PerPage).Offset((params.Page - 1) * params.PerPage).Find(&certificate).Error; err != nil {
		return nil, 0, err
	}

	return certificate, count, nil
}
