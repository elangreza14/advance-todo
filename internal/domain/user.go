package domain

//go:generate mockgen -package=domain_test -destination=./mock/mock_user_test.go github.com/elangreza14/advance-todo/internal/domain UserRepository

import (
	"context"

	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	// User is user struct
	User struct {
		ID       uuid.UUID `db:"id"`
		Email    string    `db:"email"`
		FullName string    `db:"full_name"`
		Password string    `db:"password"`

		Versioning
	}

	// UserRepository is behavior of user
	UserRepository interface {
		GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
		GetUserByEmail(ctx context.Context, email string) (*User, error)
		CreateUser(ctx context.Context, req User) (*uuid.UUID, error)
	}
)

// NewUser is new user constructor
func NewUser(req dto.RegisterUserRequest) User {
	return User{
		ID:       uuid.New(),
		Email:    req.Email,
		FullName: req.FullName,
	}
}

// SetPassword used to set the password
func (u *User) SetPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pass)
	return nil
}

// ValidatePassword used to validate the password
func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
