package model

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	CertificateID     uuid.UUID `json:"certificate_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID            string    `json:"user_id" gorm:"type:uuid;not null"`
	FileName          string    `json:"file_name" gorm:"type:varchar(255);not null"`
	FileUrl           string    `json:"file_url" gorm:"type:varchar(255);not null"`
	MD5               string    `json:"md5" gorm:"type:varchar(255);not null"`
	BlockchainAddress string    `json:"blockchain_address" gorm:"type:varchar(255);not null"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp; default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp; default:now()"`
	DeleteAt  time.Time `json:"delete_at" gorm:"type:timestamp; default:null"`

	User User `json:"user"`
}

func (Certificate) TableName() string {
	return "tr_certificates"
}
