package model

type Assigment struct {
	AssigmentID    int    `gorm:"primaryKey;autoIncrement;not null" json:"assigment_id"`
	Title          string `gorm:"type:varchar(100);not null" json:"title"`
	Description    string `gorm:"type:text" json:"description"`
	CourseDetailID int    `gorm:"type:int;not null" json:"course_detail_id"`
	BaseCode       string `gorm:"type:text" json:"base_code"`
	TestingCode    string `gorm:"type:text" json:"testing_code"`
	CreatedAt      string `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt      string `gorm:"type:timestamp;not null" json:"updated_at"`
	DeletedAt      string `gorm:"type:timestamp;null" json:"deleted_at"`
}

func (Assigment) TableName() string {
	return "tr_assigments"
}
