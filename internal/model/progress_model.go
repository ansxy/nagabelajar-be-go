package model

type Progress struct {
	ProgressID     int    `json:"progress_id" gorm:"primaryKey;autoIncrement"`
	UserID         string `json:"user_id" gorm:"type:int;not null"`
	CourseID       int    `json:"course_id" gorm:"type:int;not null"`
	CourseDetailID int    `json:"course_detail_id" gorm:"type:int;not null"`
	IsFinished     bool   `json:"is_finished" gorm:"type:boolean;not null"`
}

func (Progress) TableName() string {
	return "tr_progress"
}
