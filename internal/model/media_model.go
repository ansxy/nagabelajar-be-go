package model

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	MediaID  uuid.UUID `json:"media_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	Type     string    `gorm:"type:varchar(100);not null" json:"type"`
	UrlMedia string    `gorm:"type:text;not null" json:"url_media"`

	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updated_at"`
}

func (Media) TableName() string {
	return "tr_media"
}
