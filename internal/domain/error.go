package domain

import "errors"

var (
	ErrorEmailAlreadyExist = errors.New("user with the current email already exist")
	ErrorNotFoundEmail     = errors.New("user is not found")
	ErrorUserAndPassword   = errors.New("user or password is wrong")
)
