package main

// import (
// 	"log"

// 	"github.com/elangreza14/advance-todo/config"
// 	"github.com/elangreza14/advance-todo/internal/infrastructure"
// )

// func main() {
// 	env, err := config.NewEnv()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err = infrastructure.Run(env); err != nil {
// 		log.Fatal(err)
// 	}
// }

import (
	"fmt"
	"log"
	"time"

	"github.com/elangreza14/advance-todo/adapter/token"
)

func main() {
	tokenGen := token.NewGeneratorToken()

	res, err := tokenGen.Claims(time.Hour)
	if err != nil {
		log.Fatal("cek 1 ", err)
	}

	fmt.Println(*res)

	resp, err := tokenGen.Validate(res.Token)
	if err != nil {
		log.Fatal("cek 2 ", err)
	}
	fmt.Println(resp)
}
