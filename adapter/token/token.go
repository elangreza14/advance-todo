package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	TokenGenerator struct {
		ID        uuid.UUID
		Token     string
		ExpiredAt time.Time
		IssuedAt  time.Time
	}
	GeneratorToken interface {
		Claims(duration time.Duration) (*TokenGenerator, error)
		Validate(token string) (*TokenGenerator, error)
	}
)

var (
	ErrTokenIsExpired error = errors.New("token is expired")
	ErrParsingToken   error = errors.New("error when parsing token")
)
