package model

type CourseDetail struct {
	CourseDetailID int    `gorm:"primaryKey;autoIncrement;not null" json:"course_detail_id"`
	Name           string `gorm:"type:varchar(100);not null" json:"name"`
	Position       int    `gorm:"type:int;not null" json:"position"`
	CourseID       int    `gorm:"type:int;not null" json:"course_id"`
	Objective      string `gorm:"type:text" json:"objective"`

	Content   []CourseContent `json:"content"`
	Assigment []Assigment     `json:"assigment"`
}

func (CourseDetail) TableName() string {
	return "tr_course_details"
}
