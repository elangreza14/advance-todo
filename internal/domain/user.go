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
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pass)
	return nil
}

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
