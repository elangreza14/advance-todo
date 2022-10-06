package config

import (
	"database/sql"
)

type (
	Configuration struct {
		DbSql  *sql.DB
		Env    *Env
		Logger Logger
	}

	Option interface {
		apply(*Configuration) error
	}
)

func NewConfig(
	env *Env,
	opts ...Option,
) (*Configuration, error) {
	config := Configuration{Env: env}

	for _, o := range opts {
		if err := o.apply(&config); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
