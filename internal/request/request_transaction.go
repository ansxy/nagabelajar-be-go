package request

import "github.com/google/uuid"

type InsertTransaction struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	CourseID int       `json:"course_id" validate:"required"`
}
