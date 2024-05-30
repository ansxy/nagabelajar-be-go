package model

type CoursePractice struct {
	CoursePracticeID int    `gorm:"primaryKey;autoIncrement;not null" json:"course_practice_id"`
	CourseContentID  int    `gorm:"type:int;not null" json:"course_content_id"`
	Title            string `gorm:"type:varchar(100);not null" json:"title"`
	Content          string `gorm:"type:text;not null" json:"content"`
}

func (CoursePractice) TableName() string {
	return "tr_course_practices"
}
