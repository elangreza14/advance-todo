package auth

import (
	"context"
	"database/sql"
	"time"

	tokenAdapter "github.com/elangreza14/advance-todo/adapter/token"
	"github.com/elangreza14/advance-todo/config"
	domain "github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/internal/dto"
)

func NewAuthService(
	configuration *config.Configuration,
	authRepo domain.UserRepository,
	tokenRepo domain.TokenRepository) AuthService {

	return &authService{
		authRepo:  authRepo,
		tokenRepo: tokenRepo,
		gen:       tokenAdapter.NewGeneratorToken(configuration),
		conf:      configuration,
	}
}

func (as *authService) RegisterUser(ctx context.Context, req dto.RegisterUserRequest) error {
	if user, err := as.authRepo.GetUserByEmail(ctx, req.Email); (err != nil && err != sql.ErrNoRows) || user != nil {
		err := domain.ErrorEmailAlreadyExist
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

	if err := user.ValidatePassword(req.Password); err != nil {
		as.conf.Logger.Error("user.ValidatePassword", err)
		return nil, domain.ErrorUserAndPassword
	}

	token, err := as.tokenRepo.GetTokenByUserID(ctx, user.ID)
	if err != nil && err != sql.ErrNoRows {
		as.conf.Logger.Error("tokenRepo.GetTokenByUserID", err)
		return nil, err

	}

	if token != nil {
		tg, err := as.gen.Validate(token.Token)
		if tg != nil && err == nil {
			return &dto.LoginUserResponse{
				Token: tg.Token,
			}, nil
		}
	}

	tg, err := as.gen.Claims(1 * time.Minute)
	if err != nil {
		as.conf.Logger.Error("gen.Claims", err)
		return nil, err
	}

	token = domain.NewToken(*tg, *user)

	if _, err := as.tokenRepo.CreateToken(ctx, *token); err != nil {
		as.conf.Logger.Error("tokenRepo.CreateToken", err)
		return nil, err
	}

	return &dto.LoginUserResponse{
		Token: token.Token,
	}, nil
}
