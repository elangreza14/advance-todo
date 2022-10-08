package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	Token struct {
		ID        uuid.UUID
		UserID    uuid.UUID
		Token     string
		ExpiredAt time.Time
		IssuedAt  time.Time

		Versioning
	}

	TokenGenerator struct {
		ID        uuid.UUID
		Token     string
		ExpiredAt time.Time
		IssuedAt  time.Time
	}

	TokenRepository interface {
		GetTokenByUserID(UserID uuid.UUID) (*Token, error)
		CreateToken(UserID uuid.UUID) (*Token, error)
	}
)

func NewToken(req User) Token {
	return Token{
		UserID:     req.ID,       // user id
		Token:      "",           // generate from jwt // create interface jwt
		ExpiredAt:  time.Time{},  // based on jwt package
		Versioning: Versioning{}, // based on jwt package
	}
}

func (u *User) Validate(token string) (*Token, error) {
	return nil, nil
}
