package model

import "github.com/google/uuid"

type Transaction struct {
	TransactionID int       `gorm:"primaryKey;autoIncrement;not null" json:"transaction_id"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CourseID      int       `json:"course_id" gorm:"not null"`

	User   User   `json:"user" gorm:"foreignKey:UserID;references:UserID"`
	Course Course `json:"course" gorm:"foreignKey:CourseID;references:CourseID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
