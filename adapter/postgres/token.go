package postgres_repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/domain"
	"github.com/google/uuid"
)

type (
	iTokenRepo interface {
		GetTokenByIDAndUserID(ctx context.Context, id, userId uuid.UUID) (*domain.Token, error)
		GetTokenByUserID(ctx context.Context, userId uuid.UUID) (*domain.Token, error)
		CreateToken(ctx context.Context, req domain.Token) (*uuid.UUID, error)
	}

	tokenRepo struct {
		db     *sql.DB
		logger config.Logger
	}
)

const (
	getTokenByIDAndUserIDQuery string = `SELECT id, user_id, token, token_type, expired_at, issued_at, version, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM "tokens" WHERE id=$1 AND user_id=$2;`
	getTokenByUserIDQuery      string = `SELECT id, user_id, token, token_type, expired_at, issued_at, version, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM "tokens" WHERE user_id=$1 ORDER BY created_at DESC LIMIT 1;`
	createTokenQuery           string = `INSERT INTO tokens(id, user_id, token, token_type, expired_at, issued_at, created_by) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
)

func WithToken() PostgresOption {
	return &tokenRepo{}
}

func (t *tokenRepo) apply(configuration *config.Configuration, repo *PostgresRepo) {
	t.db = configuration.DbSql
	t.logger = configuration.Logger
	repo.Token = t
}

func (t *tokenRepo) GetTokenByIDAndUserID(ctx context.Context, id, userId uuid.UUID) (*domain.Token, error) {
	ctx, cancel := t.NewContext(ctx)
	defer cancel()

	res := &domain.Token{}
	err := t.db.QueryRowContext(ctx, getTokenByIDAndUserIDQuery, id, userId).Scan(
		&res.ID,
		&res.UserID,
		&res.Token,
		&res.TokenType,
		&res.ExpiredAt,
		&res.IssuedAt,
		&res.Version,
		&res.CreatedAt,
		&res.CreatedBy,
		&res.UpdatedAt,
		&res.UpdatedBy,
		&res.DeletedAt,
		&res.DeletedBy,
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *tokenRepo) GetTokenByUserID(ctx context.Context, userId uuid.UUID) (*domain.Token, error) {
	ctx, cancel := t.NewContext(ctx)
	defer cancel()

	res := &domain.Token{}
	err := t.db.QueryRowContext(ctx, getTokenByUserIDQuery, userId).Scan(
		&res.ID,
		&res.UserID,
		&res.Token,
		&res.TokenType,
		&res.ExpiredAt,
		&res.IssuedAt,
		&res.Version,
		&res.CreatedAt,
		&res.CreatedBy,
		&res.UpdatedAt,
		&res.UpdatedBy,
		&res.DeletedAt,
		&res.DeletedBy,
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *tokenRepo) CreateToken(ctx context.Context, req domain.Token) (*uuid.UUID, error) {
	ctx, cancel := t.NewContext(ctx)
	defer cancel()

	stmt, err := t.db.Prepare(createTokenQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res := &uuid.UUID{}
	if err := stmt.QueryRowContext(ctx,
		req.ID,
		req.UserID,
		req.Token,
		req.TokenType,
		req.ExpiredAt,
		req.IssuedAt,
		req.UserID,
	).Scan(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (t *tokenRepo) NewContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithTimeout(ctx, 10*time.Second)
}
