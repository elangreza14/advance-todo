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
	iUserRepo interface {
		GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
		GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
		CreateUser(ctx context.Context, req domain.User) (*uuid.UUID, error)
	}

	userRepo struct {
		db     *sql.DB
		logger config.Logger
	}
)

const (
	getUserByEmailQuery string = `SELECT id, email, full_name, password, version, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM "users" WHERE email=$1`
	getUserByIDQuery    string = `SELECT id, email, full_name, password, version, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM "users" WHERE id=$1`
	createUserQuery     string = `INSERT INTO users(id, email, full_name, password, created_by) VALUES($1, $2, $3, $4, $5) RETURNING id`
)

func WithUser() PostgresOption {
	return &userRepo{}
}

func (u *userRepo) apply(configuration *config.Configuration, repo *PostgresRepo) {
	u.db = configuration.DbSql
	u.logger = configuration.Logger
	repo.User = u
}

func (u *userRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	ctx, cancel := u.NewContext(ctx)
	defer cancel()

	res := &domain.User{}
	err := u.db.QueryRowContext(ctx, getUserByIDQuery, id).Scan(
		&res.ID,
		&res.Email,
		&res.FullName,
		&res.Password,
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

func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := u.NewContext(ctx)
	defer cancel()

	res := &domain.User{}
	err := u.db.QueryRowContext(ctx, getUserByEmailQuery, email).Scan(
		&res.ID,
		&res.Email,
		&res.FullName,
		&res.Password,
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

func (u *userRepo) CreateUser(ctx context.Context, req domain.User) (*uuid.UUID, error) {
	ctx, cancel := u.NewContext(ctx)
	defer cancel()

	stmt, err := u.db.Prepare(createUserQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res := &uuid.UUID{}
	if err := stmt.QueryRowContext(ctx,
		req.ID,
		req.Email,
		req.FullName,
		req.Password,
		req.ID).Scan(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userRepo) NewContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithTimeout(ctx, 10*time.Second)
}
