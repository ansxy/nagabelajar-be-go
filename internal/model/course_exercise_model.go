package model

type CourseExercise struct {
	CourseExerciseID int    `gorm:"primaryKey;autoIncrement;not null" json:"course_exercise_id"`
	Title            string `gorm:"type:varchar(100);not null" json:"title"`
	Content          string `gorm:"type:text;not null" json:"content"`
	CourseContentID  int    `gorm:"type:int;not null" json:"course_content_id"`
}

func (CourseExercise) TableName() string {
	return "tr_course_exercises"
}
