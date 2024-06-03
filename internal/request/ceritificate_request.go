package request

type CreateCertificateRequest struct {
	FirebaseID string `json:"-"`
	CourseId   int    `json:"course_id"`
}

type ListCertificateRequest struct {
	BaseQuery
	UserID     *string `json:"-"`
	FirebaseID string  `json:"-"`
}
