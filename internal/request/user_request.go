package request

type UpsertUserRequest struct {
	Name     string `json:"name" required:"true" validate:"required"`
	Email    string `json:"email" required:"true" validate:"required,email"`
	Password string `json:"password" required:"true" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" required:"true" validate:"required,email"`
	Password string `json:"password" required:"true" validate:"required"`
}
