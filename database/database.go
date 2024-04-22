package database

import (
	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	db.AutoMigrate(
		&model.User{},
		&model.UserInfo{},
		&model.Course{},
		&model.Category{},
		&model.Certificate{},
		// &model.Transaction{},
		&model.CourseContent{},
	)

	return nil
}
