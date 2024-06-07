package database

import (
	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	db.AutoMigrate(
		// &model.Transaction{},
		&model.User{},
		&model.UserInfo{},
		&model.Course{},
		&model.Category{},
		&model.Certificate{},
		&model.CourseContent{},
		&model.CourseDetail{},
		&model.CourseExercise{},
		&model.CoursePractice{},
		&model.CourseSubContent{},
		&model.Enrollment{},
		&model.Progress{},
		&model.Assigment{},
	)

	return nil
}
