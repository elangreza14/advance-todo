package domain

import (
	"context"

	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		ID       uuid.UUID `db:"id"`
		Email    string    `db:"email"`
		FullName string    `db:"full_name"`
		Password string    `db:"password"`

		Versioning
	}
)

const costBcrypt = 16

type UserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, req User) (*uuid.UUID, error)
}

func NewUser(req dto.RegisterUserRequest) User {
	return User{
		Email:    req.Email,
		FullName: req.FullName,
	}
}

func (u *User) SetPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), costBcrypt)
	if err != nil {
		return err
	}
	u.Password = string(pass)
	return nil
}
