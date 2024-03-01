package model

type CourseContent struct {
	CourseDetailID int    `gorm:"primaryKey;autoIncrement;not null" json:"course_detail_id"`
	CourseID       int    `json:"course_id" gorm:"not null"`
	Position       int    `json:"position" gorm:"not null"`
	Title          string `gorm:"type:varchar(100);not null" json:"title"`
	Content        string `gorm:"type:text;not null" json:"content"`
	VideoURL       string `gorm:"type:text;" json:"video_url"`

	Course Course `json:"course"`
}

func (CourseContent) TableName() string {
	return "tr_course_content"
}
