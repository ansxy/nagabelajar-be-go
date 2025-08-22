package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

// CreateUser implements IFaceRepository.
func (repo *Repository) CreateUser(ctx context.Context, data *model.User) error {
	return repo.BaseRepository.Create(repo.db.WithContext(ctx), data)
}

// FindOneUser implements IFaceRepository.
func (repo *Repository) FindOneUser(ctx context.Context, query ...interface{}) (*model.User, error) {
	var res *model.User

	if err := repo.BaseRepository.FindOne(
		repo.db.WithContext(ctx).Where(query[0], query[1:]...), &res,
	); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateUser implements IFaceRepository.
func (repo *Repository) UpdateUser(ctx context.Context, data *model.User) error {
	return repo.BaseRepository.Update(repo.db.WithContext(ctx), data)
}

// FindUsers implements IFaceRepository.
func (repo *Repository) FindUsers(ctx context.Context, params *request.ListUserRequest) ([]model.User, int64, error) {
	var res []model.User

	if err := repo.BaseRepository.FindList(
		repo.db.WithContext(ctx).Where("role = ?", params.Role), &res,
	); err != nil {
		return nil, 0, err
	}

	return res, 0, nil

}
