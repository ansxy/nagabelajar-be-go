package model

type CourseSubContent struct {
	CourseSubContentID int    `gorm:"primaryKey;autoIncrement;not null" json:"course_subcontent_id"`
	Title              string `gorm:"type:varchar(100);not null" json:"title"`
	Content            string `gorm:"type:text;not null" json:"content"`
	CourseContentID    int    `gorm:"type:int;not null" json:"course_content_id"`
}

func (CourseSubContent) TableName() string {
	return "tr_course_subcontents"
}
