package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository struct{}

// Parameter Data is Model struct
func (repo *BaseRepository) Create(db *gorm.DB, data interface{}) error {
	err := db.Create(data).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *BaseRepository) Update(db *gorm.DB, data interface{}) error {
	err := db.Model(data).Updates(data).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *BaseRepository) Delete(db *gorm.DB, data interface{}) error {
	err := db.Delete(data).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *BaseRepository) FindOne(db *gorm.DB, result interface{}) (err error) {
	err = db.First(result).Preload(clause.Associations).Error
	return
}
