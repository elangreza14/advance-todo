package dto

type (
	RegisterUserRequest struct {
		Email    string `json:"email" validate:"required, email"`
		FullName string `json:"full_name" validate:"required, gte=6"`
		Password string `json:"password"  validate:"required, gte=6"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password"  validate:"required, gte=6"`
	}

	LoginUserResponse struct {
		Token string `json:"token" validate:"required, email"`
	}
)
