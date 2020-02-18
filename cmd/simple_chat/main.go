package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

var (
	port = "8080"
)

func main() {
	log.Printf("Listening %s\n", port)

	// global context
	ctx, cancel := context.WithCancel(context.Background())

	// handle interrupt signal
	// call `cancel` function on interrupt signal
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		select {
		case <-signalChannel:
			log.Println("Closing...")
			cancel()
		}
	}()

	// TODO: connect to repository, etc ....

	<-ctx.Done() // lock finishing `main` function until got interrupt signal
}
