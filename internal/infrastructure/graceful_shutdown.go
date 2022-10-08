package infrastructure

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	Job         func(context.Context) error
	JobFunction map[string]Job
	resultJob   struct {
		name string
		err  error
	}
)

func GracefulShutdown(job JobFunction) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-c

	fmt.Print("\nrunning cleanup function...")

	errors := []error{}
	functionResult := make(chan resultJob, len(job))

	for i, v := range job {
		go func(functionName string, vf Job) {

			fmt.Printf("\nstart cleanup %v", functionName)

			ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
			defer cancel()

			if err := vf(ctx); err != nil {
				errors = append(errors, err)
				functionResult <- resultJob{
					name: functionName,
					err:  err,
				}
			} else {
				functionResult <- resultJob{
					name: functionName,
					err:  nil,
				}
			}

		}(i, v)
	}

	wait := time.After(time.Duration(len(job)*10) * time.Second)

	for range job {
		select {
		case fr := <-functionResult:
			if fr.err != nil {
				fmt.Printf("\nerror when cleanup %v: %v", fr.name, fr.err)
			} else {
				fmt.Printf("\nfinish cleanup %v", fr.name)
			}
		case <-wait:
			fmt.Println("\ntimeout")
			return
		}
	}

	fmt.Println("\napplication successfully shutdown")

}
