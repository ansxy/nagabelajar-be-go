package model

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	EnrollmentID int    `gorm:"primaryKey;autoIncrement;not null" json:"enrollment_id"`
	UserID       string `gorm:"type:varchar(100);not null" json:"user_id"`
	CourseID     string `gorm:"type:varchar(100);not null" json:"course_id"`

	CreatedAt time.Time      `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"-"`
}

func (Enrollment) TableName() string {
	return "tr_enrollments"
}
