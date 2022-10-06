package core

import (
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	postgresRepo "github.com/elangreza14/advance-todo/adapter/postgres"

)

type Core struct {
	AuthService auth.AuthService
}

func New(conf *config.Configuration, pr *postgresRepo.PostgresRepo) Core {
	authService := auth.NewAuthService(conf, pr.User)
	return Core{
		AuthService: authService,
	}
}
