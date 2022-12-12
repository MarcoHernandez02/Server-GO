package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8000")

	go srv.ListenAndServe()
	log.Println("Server started")

	<-serverDoneChan

	srv.shutdown(ctx)
	log.Println("Server stopped")

}
