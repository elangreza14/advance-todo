package auth

import (
	"context"
	"database/sql"

	"github.com/elangreza14/advance-todo/config"
	domain "github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/dto"
)

func NewAuthService(
	configuration *config.Configuration,
	authRepo domain.UserRepository) AuthService {
	return &authService{
		authRepo: authRepo,
		conf:     configuration,
	}
}

func (as *authService) RegisterUser(ctx context.Context, req dto.RegisterUserRequest) error {
	if user, err := as.authRepo.GetUserByEmail(ctx, req.Email); (err != nil && err != sql.ErrNoRows) || user != nil {
		err := domain.ErrorNotFoundEmail
		as.conf.Logger.Error("authRepo.GetUserByEmail", err)
		return err
	}

	user := domain.NewUser(req)
	if err := user.SetPassword(req.Password); err != nil {
		as.conf.Logger.Error("user.SetPassword", err)
		return err
	}

	_, err := as.authRepo.CreateUser(ctx, user)
	if err != nil {
		as.conf.Logger.Error("authRepo.CreateUser", err)
		return err
	}

	return nil
}

func (as *authService) LoginUser(ctx context.Context, req dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	user, err := as.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			as.conf.Logger.Error("authRepo.GetUserByEmail.sql.ErrNoRows", err)
			return nil, domain.ErrorNotFoundEmail
		}
		as.conf.Logger.Error("authRepo.GetUserByEmail", err)
		return nil, err
	}

	if user.Password != req.Password {
		as.conf.Logger.Error("user.Password != req.Password", err)
		return nil, domain.ErrorNotFoundEmail
	}

	return &dto.LoginUserResponse{
		Token: "",
	}, nil
}
