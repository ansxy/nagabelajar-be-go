package model

import "time"

type Category struct {
	CategoryID int       `gorm:"primaryKey;" json:"category_id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp; default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp; default:now()"`
	DeletedAt  time.Time `json:"-" gorm:"type:timestamp; default:null"`

	Courses []Course `json:"courses"`
}

func (Category) TableName() string {
	return "tm_category"
}
