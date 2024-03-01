package repository

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

//CreeateCategory implements IFaceRepository.

func (repo *Repository) CreateCategory(ctx context.Context, data *model.Category) error {
	return repo.BaseRepository.Create(repo.db.WithContext(ctx), data)
}

// FindListCategory implements IFaceRepository.
func (repo *Repository) FindListCategory(ctx context.Context, params *request.ListCategoryRequest) ([]model.Category, int64, error) {
	var category []model.Category
	var count int64

	query := repo.db.WithContext(ctx).Model(&model.Category{})

	if params.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+params.Keyword+"%")
	}

	if params.Sort != "" {
		query = query.Order(params.Sort)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(params.PerPage).Offset((params.Page - 1) * params.PerPage).Find(&category).Error; err != nil {
		return nil, 0, err
	}

	return category, count, nil

}

// DeleteOneCategory implements IFaceRepository.
func (repo *Repository) DeleteOneCategory(ctx context.Context, categoryID int) error {
	return repo.BaseRepository.Delete(repo.db, &model.Category{CategoryID: categoryID})
}

// FindOneCategory implements IFaceRepository.
func (repo *Repository) FindOneCategory(ctx context.Context, categoryID int) (*model.Category, error) {
	var res *model.Category

	if err := repo.BaseRepository.FindOne(repo.db.WithContext(ctx).Where("category_id = ?", categoryID), &res); err != nil {
		return nil, err
	}

	return res, nil
}
