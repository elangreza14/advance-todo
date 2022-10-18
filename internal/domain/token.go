package domain

import (
	"context"
	"time"

	"github.com/elangreza14/advance-todo/config"

	"github.com/google/uuid"
)

type (
	TokenType string
	Token     struct {
		ID        uuid.UUID
		UserID    uuid.UUID
		Token     string
		IP        string
		TokenType TokenType
		ExpiredAt time.Time
		IssuedAt  time.Time

		Versioning
	}

	TokenRepository interface {
		GetTokenByIDAndUserID(ctx context.Context, id, userId uuid.UUID) (*Token, error)
		GetTokenByUserIDAndIP(ctx context.Context, userId uuid.UUID, ip string, tokenType TokenType) (*Token, error)
		CreateToken(ctx context.Context, req Token) (*uuid.UUID, error)
	}
)

const (
	TokenTypePassword TokenType = "PASSWORD"
	TokenTypeAccess   TokenType = "ACCESS"
	TokenTypeRefresh  TokenType = "REFRESH"
)

func NewToken(gen config.TokenGenerator, req User, tokenType TokenType, ip string) *Token {
	return &Token{
		ID:        gen.ID,
		UserID:    req.ID,
		Token:     gen.Token,
		ExpiredAt: gen.ExpiredAt,
		IssuedAt:  gen.IssuedAt,
		TokenType: tokenType,
		IP:        ip,
	}
}
