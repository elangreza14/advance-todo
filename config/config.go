package config

import (
	"database/sql"
)

type (
	// Configuration is struct handle all configuration
	Configuration struct {
		DbSQL     *sql.DB
		Env       *Env
		Logger    ILogger
		Cache     ICache
		Token     IToken
		Validator IValidator
	}

	// Option is interface to grouping all configuration
	Option interface {
		apply(*Configuration) error
	}
)

// NewConfig is creating new configuration
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
