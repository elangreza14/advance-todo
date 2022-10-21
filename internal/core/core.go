package core

import (
	postgresRepo "github.com/elangreza14/advance-todo/adapter/postgres"
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
)

type Core struct {
	AuthService auth.AuthService
}

func New(conf *config.Configuration, pr *postgresRepo.PostgresRepo) Core {
	authService := auth.NewAuthService(conf, pr.User, pr.Token)
	return Core{
		AuthService: authService,
	}
}
