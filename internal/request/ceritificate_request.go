package request

import "github.com/google/uuid"

type CreateCertificateRequest struct {
	UserId   uuid.UUID `json:"-"`
	CourseId int       `json:"course_id"`
}
