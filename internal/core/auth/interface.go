package auth

//go:generate mockgen -package=auth_test -destination=./mock_auth_test.go github.com/elangreza14/advance-todo/internal/core/auth AuthService

import (
	"context"

	"github.com/elangreza14/advance-todo/config"
	domain "github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/internal/dto"
)

type (
	AuthService interface {
		RegisterUser(ctx context.Context, req dto.RegisterUserRequest) error
		LoginUser(ctx context.Context, req dto.LoginUserRequest) (*dto.LoginUserResponse, error)
		GetUser(ctx context.Context) (*dto.UserDetailResponse, error)
	}

	authService struct {
		authRepo  domain.UserRepository
		tokenRepo domain.TokenRepository
		conf      *config.Configuration
	}
)
