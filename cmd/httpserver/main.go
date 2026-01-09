package main

import (
	"httpserver/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = 42069

func main() {
	s, err := server.Serve(port)
	if err != nil {
		log.Fatal("Error starting server: %v", err)
	}

	defer s.Close()
	log.Println("Server started on port", port)
	// new channel that can only carry values of type os.Signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("\nServer gracefully stopped")
}
