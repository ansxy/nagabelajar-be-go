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

type LoginWithGoogleRequest struct {
	Email      string `json:"email" required:"true" validate:"required,email"`
	Name       string `json:"name" required:"true" validate:"required"`
	FirebaseID string `json:"-"`
}

type ListUserRequest struct {
	BaseQuery
	Role string `json:"role"`
}
