package request

type UpsertCategoryRequest struct {
	Name string `json:"name" required:"true" validate:"required"`
}

type ListCategoryRequest struct {
	BaseQuery
}
