package model

import "github.com/google/uuid"

type Transaction struct {
	TransactionID uuid.UUID `json:"transaction_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID        string    `json:"user_id" gorm:"type:uuid;not null"`
	CourseID      string    `json:"course_id" gorm:"not null"`

	User   User   `json:"user" gorm:"foreignKey:UserID;references:UserID"`
	Course Course `json:"course" gorm:"foreignKey:CourseID;references:CourseID"`
}

func (Transaction) TableName() string {
	return "tr_transactions"
}
