package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

// CreateCategory implements IFaceUsecase.
func (u *Usecase) CreateCategory(ctx context.Context, data *request.UpsertCategoryRequest) error {
	category := &model.Category{
		Name: data.Name,
	}
	return u.Repo.CreateCategory(ctx, category)
}

// DeleteOneCategory implements IFaceUsecase.
func (u *Usecase) DeleteOneCategory(ctx context.Context, categoryID int) error {
	return u.Repo.DeleteOneCategory(ctx, categoryID)
}

// FindListCategory implements IFaceUsecase.
func (u *Usecase) FindListCategory(ctx context.Context, params *request.ListCategoryRequest) ([]model.Category, int64, error) {
	res, cnt, err := u.Repo.FindListCategory(ctx, params)
	return res, cnt, err
}

// FindOneCategory implements IFaceUsecase.
func (u *Usecase) FindOneCategory(ctx context.Context, categoryID int) (*model.Category, error) {
	return u.Repo.FindOneCategory(ctx, categoryID)
}
