package domain

import "errors"

var (
	ErrorNotFoundEmail = errors.New("user with the current email already exist")
	ErrorUserAndPassword = errors.New("user or password is wrong")
)
