package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Certificate struct {
	CertificateID     uuid.UUID `json:"certificate_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID            string    `json:"user_id" gorm:"type:uuid;not null"`
	CourseID          string    `json:"course_id" gorm:"type:uuid;"`
	FileName          string    `json:"file_name" gorm:"type:varchar(255);not null"`
	FileUrl           string    `json:"file_url" gorm:"type:varchar(255);not null"`
	MD5               string    `json:"md5" gorm:"type:varchar(255);not null"`
	BlockchainAddress string    `json:"blockchain_address" gorm:"type:varchar(255);not null"`
	GasUsed           string    `json:"gas_used" gorm:"type:numeric;"`
	GasPrice          string    `json:"gas_price" gorm:"type:numeric;"`
	Cost              string    `json:"cost" gorm:"type:numeric;"`

	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp; default:now()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp; default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"-"`
	Course    Course         `json:"course"`
}

func (Certificate) TableName() string {
	return "tr_certificates"
}
