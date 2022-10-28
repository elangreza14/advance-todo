package auth

//go:generate mockgen -package=auth_test -destination=./mock_auth_test.go github.com/elangreza14/advance-todo/internal/core/auth AuthService

import (
	"context"

	"github.com/elangreza14/advance-todo/config"
	domain "github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/google/uuid"
)

type (
	AuthService interface {
		RegisterUser(ctx context.Context, req dto.RegisterUserRequest) error
		LoginUser(ctx context.Context, req dto.LoginUserRequest) (*dto.LoginUserResponse, error)
		GetUser(ctx context.Context) (*dto.UserDetailResponse, error)
		GetTokenByID(ctx context.Context, id uuid.UUID) (*domain.Token, error) // use for getting header
		RefreshToken(ctx context.Context) (*dto.LoginUserResponse, error)
	}

	authService struct {
		authRepo  domain.UserRepository
		tokenRepo domain.TokenRepository
		conf      *config.Configuration
	}
)
