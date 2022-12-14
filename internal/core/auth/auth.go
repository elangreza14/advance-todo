package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/elangreza14/advance-todo/config"
	domain "github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/google/uuid"
)

// NewAuthService is a new constructor service
func NewAuthService(
	configuration *config.Configuration,
	authRepo domain.UserRepository,
	tokenRepo domain.TokenRepository,
) AuthService {
	return &authService{
		authRepo:  authRepo,
		tokenRepo: tokenRepo,
		conf:      configuration,
	}
}

func (as *authService) RegisterUser(ctx context.Context, req dto.RegisterUserRequest) error {
	if user, err := as.authRepo.GetUserByEmail(ctx, req.Email); (err != nil && err != sql.ErrNoRows) || user != nil {
		err := domain.ErrorEmailAlreadyExist
		as.conf.Logger.Error("authService.authRepo.GetUserByEmail", err)
		return err
	}

	user := domain.NewUser(req)
	if err := user.SetPassword(req.Password); err != nil {
		as.conf.Logger.Error("authService.user.SetPassword", err)
		return err
	}

	_, err := as.authRepo.CreateUser(ctx, user)
	if err != nil {
		as.conf.Logger.Error("authService.authRepo.CreateUser", err)
		return err
	}

	return nil
}

func (as *authService) LoginUser(ctx context.Context, req dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	user, err := as.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			as.conf.Logger.Error("authService.authRepo.GetUserByEmail.sql.ErrNoRows", err)
			return nil, domain.ErrorNotFoundEmail
		}
		as.conf.Logger.Error("authService.authRepo.GetUserByEmail", err)
		return nil, err
	}

	if err := user.ValidatePassword(req.Password); err != nil {
		as.conf.Logger.Error("authService.user.ValidatePassword", err)
		return nil, domain.ErrorUserAndPassword
	}

	// getting ip user from context
	ipUser := ctx.Value(domain.ContextValueIP).(string)

	tknAccess, err := as.tokenRepo.GetTokenByUserIDAndIP(ctx, user.ID, ipUser, domain.TokenTypeAccess)
	if err != nil && err != sql.ErrNoRows {
		as.conf.Logger.Error("authService.tokenRepo.GetTokenByUserIDAndIP", err)
		return nil, err
	}

	tknRefresh, err := as.tokenRepo.GetTokenByUserIDAndIP(ctx, user.ID, ipUser, domain.TokenTypeRefresh)
	if err != nil && err != sql.ErrNoRows {
		as.conf.Logger.Error("authService.tokenRepo.GetTokenByUserIDAndIP", err)
		return nil, err
	}

	if tknAccess != nil && tknRefresh != nil {
		tgAccess, errAccess := as.conf.Token.Validate(tknAccess.Token)
		tgRefresh, errRefresh := as.conf.Token.Validate(tknRefresh.Token)
		if (tgAccess != nil && errAccess == nil) &&
			(tgRefresh != nil && errRefresh == nil) {
			// if token is exist
			// just return
			return &dto.LoginUserResponse{
				AccessToken:  tgAccess.Token,
				RefreshToken: tgRefresh.Token,
			}, nil
		}
	}

	tknAccess, err = as.generateToken(ctx, domain.TokenTypeAccess, user.ID, ipUser)
	if err != nil {
		as.conf.Logger.Error("authService.generateToken", err)
		return nil, err
	}

	tknRefresh, err = as.generateToken(ctx, domain.TokenTypeRefresh, user.ID, ipUser)
	if err != nil {
		as.conf.Logger.Error("authService.generateToken", err)
		return nil, err
	}

	return &dto.LoginUserResponse{
		AccessToken:  tknAccess.Token,
		RefreshToken: tknRefresh.Token,
	}, nil
}

func (as *authService) GetUser(ctx context.Context) (*dto.UserDetailResponse, error) {
	rawUserID := ctx.Value(domain.ContextValueUserID).(string)
	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		as.conf.Logger.Error("authService.uuid.Parse", err)
		return nil, err
	}

	user, err := as.authRepo.GetUserByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			as.conf.Logger.Error("authService.authRepo.GetUserByID.sql.ErrNoRows", err)
			return nil, domain.ErrorNotFoundEmail
		}
		as.conf.Logger.Error("authService.authRepo.GetUserByID", err)
		return nil, err
	}

	return &dto.UserDetailResponse{
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (as *authService) GetTokenByID(ctx context.Context, id uuid.UUID) (*domain.Token, error) {
	res, err := as.tokenRepo.GetTokenByID(ctx, id)
	if err != nil {
		as.conf.Logger.Error("authService.tokenRepo.GetTokenByID", err)
		return nil, err
	}

	return res, nil
}

func (as *authService) RefreshToken(ctx context.Context) (*dto.LoginUserResponse, error) {
	ipUser := ctx.Value(domain.ContextValueIP).(string)
	rawUserID := ctx.Value(domain.ContextValueUserID).(string)
	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		as.conf.Logger.Error("authService.uuid.Parse", err)
		return nil, err
	}

	tknAccess, err := as.tokenRepo.GetTokenByUserIDAndIP(ctx, userID, ipUser, domain.TokenTypeAccess)
	if err != nil && err != sql.ErrNoRows {
		as.conf.Logger.Error("authService.tokenRepo.GetTokenByUserIDAndIP", err)
		return nil, err
	}

	tknRefresh, err := as.tokenRepo.GetTokenByUserIDAndIP(ctx, userID, ipUser, domain.TokenTypeRefresh)
	if err != nil && err != sql.ErrNoRows {
		as.conf.Logger.Error("authService.tokenRepo.GetTokenByUserIDAndIP", err)
		return nil, err
	}

	if tknAccess != nil && tknRefresh != nil {
		tgAccess, errAccess := as.conf.Token.Validate(tknAccess.Token)
		tgRefresh, errRefresh := as.conf.Token.Validate(tknRefresh.Token)
		if (tgAccess != nil && errAccess == nil) &&
			(tgRefresh != nil && errRefresh == nil) {
			// if token is exist
			// just return
			return &dto.LoginUserResponse{
				AccessToken:  tgAccess.Token,
				RefreshToken: tgRefresh.Token,
			}, nil
		}
	}

	tknAccess, err = as.generateToken(ctx, domain.TokenTypeAccess, userID, ipUser)
	if err != nil {
		as.conf.Logger.Error("authService.generateToken", err)
		return nil, err
	}

	tknRefresh, err = as.generateToken(ctx, domain.TokenTypeRefresh, userID, ipUser)
	if err != nil {
		as.conf.Logger.Error("authService.generateToken", err)
		return nil, err
	}

	return &dto.LoginUserResponse{
		AccessToken:  tknAccess.Token,
		RefreshToken: tknRefresh.Token,
	}, nil
}

func (as *authService) generateToken(ctx context.Context, tokenType domain.TokenType, userId uuid.UUID, ipUser string) (*domain.Token, error) {
	var duration time.Duration
	switch tokenType {
	case domain.TokenTypeRefresh:
		duration = time.Minute * 5
	case domain.TokenTypeAccess:
		duration = time.Minute * 2
	case domain.TokenTypePassword:
		duration = time.Minute * 1
	}

	tgClaim, err := as.conf.Token.Claims(duration)
	if err != nil {
		as.conf.Logger.Error("authService.conf.Token.Claims", err)
		return nil, err
	}

	res := domain.NewToken(*tgClaim, userId, tokenType, ipUser, duration)

	if _, err := as.tokenRepo.CreateToken(ctx, *res); err != nil {
		as.conf.Logger.Error("authService.tokenRepo.CreateToken", err)
		return nil, err
	}

	return res, nil
}
