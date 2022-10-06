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

	if err = infrastructure.WithApi(env); err != nil {
		log.Fatal(err)
	}
}
