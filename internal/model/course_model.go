package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	CourseID    int       `gorm:"primaryKey;autoIncrement;not null" json:"course_id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Code        string    `gorm:"type:varchar(100);not null" json:"code"`
	Author      string    `gorm:"type:varchar(100)" json:"author"`
	CategoryID  int       `gorm:"type:int;not null" json:"category_id"`
	Price       int       `gorm:"type:int;not null" json:"price"`
	Description string    `gorm:"type:text;not null" json:"description"`
	MediaID     uuid.UUID `json:"media_id" gorm:"type:uuid;not null"`
	IsPaid      *bool     `gorm:"type:boolean;default:false" json:"is_paid"`
	IsArchived  *bool     `gorm:"type:boolean;default:false" json:"is_archived"`

	CreatedAt time.Time      `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"-"`

	Category     Category       `json:"category"`
	Media        Media          `json:"media"`
	CourseDetail []CourseDetail `json:"course_detail"`
}

func (Course) TableName() string {
	return "tr_courses"
}
