package request

type ListCourseRequest struct {
	BaseQuery
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}

type UpsertCourseRequest struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"_"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsPaid      *bool  `json:"is_paid"`
	IsArchived  *bool  `json:"is_archived"`
}
