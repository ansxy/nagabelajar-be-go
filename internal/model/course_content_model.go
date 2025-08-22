package model

type CourseContent struct {
	CourseContentID int    `gorm:"primaryKey;autoIncrement;not null" json:"course_content_id"`
	CourseDetailID  int    `gorm:"type:int;not null" json:"course_detail_id"`
	Title           string `gorm:"type:varchar(100);not null" json:"title"`
	CourseContent   string `gorm:"type:text;not null" json:"course_content"`

	// Foreign key relationships
	SubContent []CourseSubContent `gorm:"foreignKey:CourseContentID;references:CourseContentID" json:"sub_content"`
	Practice   []CoursePractice   `gorm:"foreignKey:CourseContentID;references:CourseContentID" json:"practice"`
	Exercise   []CourseExercise   `gorm:"foreignKey:CourseContentID;references:CourseContentID" json:"exercise"`
}

func (CourseContent) TableName() string {
	return "tr_course_contents"
}
