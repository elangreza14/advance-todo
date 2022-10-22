package domain

import (
	"context"
	"time"

	"github.com/elangreza14/advance-todo/config"

	"github.com/google/uuid"
)

type (
	// TokenType is kind of token to save in db and cache
	TokenType string

	// TokenType is domain for handling all token
	Token struct {
		ID        uuid.UUID
		UserID    uuid.UUID
		Token     string
		IP        string
		TokenType TokenType
		ExpiredAt time.Time
		IssuedAt  time.Time

		Versioning
	}

	// TokenRepository consist all token behavior of Token
	TokenRepository interface {
		GetTokenByID(ctx context.Context, id uuid.UUID) (*Token, error)
		GetTokenByUserIDAndIP(ctx context.Context, userId uuid.UUID, ip string, tokenType TokenType) (*Token, error)
		CreateToken(ctx context.Context, req Token) (*uuid.UUID, error)
	}
)

const (
	// TokenTypePassword is type token for changing password
	TokenTypePassword TokenType = "PASSWORD"

	// TokenTypeAccess is type token for accessing API
	TokenTypeAccess TokenType = "ACCESS"

	// TokenTypeRefresh is type token for refreshing Token API
	TokenTypeRefresh TokenType = "REFRESH"
)

// NewToken is constructor for Token
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
