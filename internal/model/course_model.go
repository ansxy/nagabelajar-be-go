package model

import (
	"time"
)

type Course struct {
	CourseID    int        `gorm:"primaryKey;autoIncrement;not null" json:"course_id"`
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Code        string     `gorm:"type:varchar(100);not null" json:"code"`
	CategoryID  int        `gorm:"type:int;not null" json:"category_id"`
	Price       int        `gorm:"type:int;not null" json:"price"`
	Description string     `gorm:"type:text;not null" json:"description"`
	IsPaid      *bool      `gorm:"type:boolean;default:false" json:"is_paid"`
	IsArchived  *bool      `gorm:"type:boolean;default:false" json:"is_archived"`
	CreatedAt   time.Time  `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;default:now()" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:timestamp" json:"deleted_at"`

	Category     Category        `json:"category"`
	CourseDetail []CourseContent `json:"course_content"`
}

func (Course) TableName() string {
	return "tr_courses"
}
