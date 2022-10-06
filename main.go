package main

import (
	"log"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/infrastructure"
)

func main() {
	env, err := config.NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	infrastructure.WithApi(env)
}
