package model

import "github.com/google/uuid"

type UserInfo struct {
	UserInfoID int       `gorm:"primaryKey;autoIncrement;not null" json:"user_info_id"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Phone      string    `json:"phone" gorm:"type:varchar(100);not null"`
	Address    string    `json:"address" gorm:"type:varchar(100);not null"`
	CreateAt   string    `json:"create_at" gorm:"type:timestamp;default:now()"`
	UpdateAt   string    `json:"update_at" gorm:"type:timestamp;default:now()"`
}

func (UserInfo) TableName() string {
	return "tr_user_info"
}
