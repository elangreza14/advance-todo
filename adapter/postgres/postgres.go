package postgres_repo

import "github.com/elangreza14/advance-todo/config"

type (
	postgresRepo struct {
		User iUserRepo
	}

	PostgresOption interface {
		apply(configuration *config.Configuration, repo *postgresRepo)
	}
)

func NewPostgresRepo(conf *config.Configuration, opts ...PostgresOption) *postgresRepo {
	postgresRepo := &postgresRepo{}

	for _, opt := range opts {
		opt.apply(conf, postgresRepo)
	}

	return postgresRepo
}
