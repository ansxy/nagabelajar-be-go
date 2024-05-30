package model

import "github.com/google/uuid"

type User struct {
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	IsGoogle   bool      `json:"is_google" gorm:"type:boolean;default:false"`
	Name       string    `json:"name" gorm:"type:varchar(100);not null"`
	Email      string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password   string    `json:"-" gorm:"type:varchar(100);not null"`
	FirebaseID string    `json:"firebase_id" gorm:"type:varchar(100);not null"`
	Role       string    `json:"role" gorm:"type:varchar(100);not null"`

	CreateAt string `json:"create_at" gorm:"type:timestamp;default:now()"`
	UpdateAt string `json:"update_at" gorm:"type:timestamp;default:now()"`
	DeleteAt string `json:"delete_at" gorm:"type:timestamp;default:now()"`

	UserInfo    UserInfo      `json:"user_info" gorm:"foreignKey:UserID;references:UserID"`
	Certificate []Certificate `json:"certificate"`
}

func (User) TableName() string {
	return "tr_users"
}
