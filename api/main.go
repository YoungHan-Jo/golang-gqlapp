package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/younghan-jo/gqlgen-todos/api/server"
)

func main() {

	svc := server.NewServer()
	go func() {
		if err := svc.Run(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	svc.Shutdown()
}
