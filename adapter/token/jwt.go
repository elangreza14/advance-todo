package token

import (
	"fmt"
	"time"

	"github.com/elangreza14/advance-todo/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type (
	JWTGeneratorToken struct {
		conf *config.Configuration
	}

	CustomClaims struct {
		id uuid.UUID
		jwt.RegisteredClaims
	}
)

func NewGeneratorToken(
	conf *config.Configuration,
) GeneratorToken {
	return &JWTGeneratorToken{conf: conf}
}

func (j *JWTGeneratorToken) Claims(duration time.Duration) (*TokenGenerator, error) {
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

	res, err := token.SignedString([]byte(j.conf.Env.TOKEN_KEY))
	if err != nil {
		j.conf.Logger.Error("token.SignedString", err)
		return nil, err
	}

	return &TokenGenerator{
		ID:        id,
		Token:     res,
		ExpiredAt: exp,
		IssuedAt:  now,
	}, nil
}

func (j *JWTGeneratorToken) Validate(token string) (*TokenGenerator, error) {
	parsed, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			j.conf.Logger.Error("token.Method.(*jwt.SigningMethodHMAC)", err)
			return nil, err
		}

		return []byte(j.conf.Env.TOKEN_KEY), nil
	})

	if err != nil {
		j.conf.Logger.Error("jwt.ParseWithClaims", ErrTokenIsExpired)
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
		j.conf.Logger.Error("parsed.Claims.(*CustomClaims)", ErrTokenIsExpired)
		return nil, ErrParsingToken
	}
}
