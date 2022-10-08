package token

import (
	"time"

	"github.com/elangreza14/advance-todo/internal/domain"
)

type GeneratorToken interface {
	Claims(duration time.Duration) (*domain.TokenGenerator, error)
	Validate(token string) (*domain.TokenGenerator, error)
}
