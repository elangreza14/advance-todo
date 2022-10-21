package dto

import "time"

type (
	// RegisterUserRequest is a request register
	RegisterUserRequest struct {
		Email    string `json:"email" validate:"required, email"`
		FullName string `json:"full_name" validate:"required, gte=6"`
		Password string `json:"password"  validate:"required, gte=6"`
	}

	// LoginUserRequest is a request for login
	LoginUserRequest struct {
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password"  validate:"required, gte=6"`
	}

	// LoginUserResponse is a response to get token data
	LoginUserResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	// UserDetailResponse is a response to get user profile data
	UserDetailResponse struct {
		Email     string    `json:"email"`
		FullName  string    `json:"full_name"`
		CreatedAt time.Time `json:"created_at"`
	}
)
