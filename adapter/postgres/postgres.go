package postgres_repo

import "github.com/elangreza14/advance-todo/config"

type (
	PostgresRepo struct {
		User iUserRepo
	}

	PostgresOption interface {
		apply(configuration *config.Configuration, repo *PostgresRepo)
	}
)

func NewPostgresRepo(conf *config.Configuration, opts ...PostgresOption) *PostgresRepo {
	postgresRepo := &PostgresRepo{}

	for _, opt := range opts {
		opt.apply(conf, postgresRepo)
	}

	return postgresRepo
}
