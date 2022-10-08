package token

import (
	"fmt"
	"time"

	"github.com/elangreza14/advance-todo/internal/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type (
	JWTGeneratorToken struct{}

	CustomClaims struct {
		id uuid.UUID
		jwt.RegisteredClaims
	}
)

const signedKey = "dOag7RwL7Lr9vdP1lbav36HMdb8QyaP2"

func NewGeneratorToken() GeneratorToken {
	return &JWTGeneratorToken{}
}

func (j *JWTGeneratorToken) Claims(duration time.Duration) (*domain.TokenGenerator, error) {
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

	res, err := token.SignedString([]byte(signedKey))
	if err != nil {
		return nil, err
	}

	return &domain.TokenGenerator{
		ID:        id,
		Token:     res,
		ExpiredAt: exp,
		IssuedAt:  now,
	}, nil
}

func (j *JWTGeneratorToken) Validate(token string) (*domain.TokenGenerator, error) {
	parsed, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signedKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsed.Claims.(*CustomClaims); ok && parsed.Valid {
		return &domain.TokenGenerator{
			ID:        claims.id,
			Token:     token,
			ExpiredAt: claims.ExpiresAt.Time,
			IssuedAt:  claims.IssuedAt.Time,
		}, nil
	} else {
		return nil, fmt.Errorf("token is not valid")
	}

}
