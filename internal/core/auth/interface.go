package auth

import (
	"context"

	"github.com/elangreza14/advance-todo/config"
	domain "github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/dto"
)

type (
	AuthService interface {
		RegisterUser(ctx context.Context, req dto.RegisterUserRequest) error
		LoginUser(ctx context.Context, req dto.LoginUserRequest) (*dto.LoginUserResponse, error)
	}

	authService struct {
		authRepo domain.UserRepository
		conf     *config.Configuration
	}
)
