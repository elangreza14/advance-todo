package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type (
	TokenGenerator struct {
		ID        uuid.UUID
		Token     string
		ExpiredAt time.Time
		IssuedAt  time.Time
	}

	IToken interface {
		apply(*Configuration) error
		Claims(duration time.Duration) (*TokenGenerator, error)
		Validate(token string) (*TokenGenerator, error)
	}

	iToken struct {
		conf *Configuration
	}

	CustomClaims struct {
		id uuid.UUID
		jwt.RegisteredClaims
	}
)

var (
	ErrCreatingToken  error = errors.New("error creating token")
	ErrTokenIsExpired error = errors.New("token is expired")
	ErrParsingToken   error = errors.New("error parsing token")
)

func newToken() IToken {
	return &iToken{}
}

func WithToken() Option {
	return newToken()
}

func (it *iToken) apply(conf *Configuration) error {
	it.conf = conf
	conf.Token = it
	return nil
}

func (it *iToken) Claims(duration time.Duration) (*TokenGenerator, error) {
	id := uuid.New()
	now := time.Now()
	exp := now.Add(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        id.String(),
		},
	})

	res, err := token.SignedString([]byte(it.conf.Env.TOKEN_KEY))
	if err != nil {
		return nil, ErrCreatingToken
	}

	return &TokenGenerator{
		ID:        id,
		Token:     res,
		ExpiredAt: exp,
		IssuedAt:  now,
	}, nil
}

func (it *iToken) Validate(token string) (*TokenGenerator, error) {
	parsed, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, err
		}

		return []byte(it.conf.Env.TOKEN_KEY), nil
	})

	if err != nil {
		// error range within jwt Standard Claim validation errors
		// so we handle >= 8 with
		// ValidationErrorExpired, ValidationErrorIssuedAt, ValidationErrorId
		if err.(*jwt.ValidationError).Errors >= 8 {
			return nil, ErrTokenIsExpired
		} else {
			return nil, ErrParsingToken
		}
	}

	if claims, ok := parsed.Claims.(*CustomClaims); ok && parsed.Valid {
		return &TokenGenerator{
			ID:        claims.id,
			Token:     token,
			ExpiredAt: claims.ExpiresAt.Time,
			IssuedAt:  claims.IssuedAt.Time,
		}, nil
	} else {
		return nil, ErrParsingToken
	}
}
