package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	repo "github.com/vladimir-kopaliani/simple_chat/internal/repository"
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

	// repository
	messagesRepo, err := repo.New(ctx, repo.Configuration{
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer messagesRepo.Close()

	<-ctx.Done() // lock finishing `main` function until got interrupt signal
}
