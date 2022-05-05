package main

import (
	"context"
	"errors"
	"go-gqlgen/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Define graphql handlers
	gh := handler.NewGraphqlHandler()

	// Define servemux
	m := http.NewServeMux()

	// Enable playground route in development
	if os.Getenv("GO_ENV") == "dev" {
		m.Handle("/", gh.Playground())
	}

	m.Handle("/query", gh.Query())

	// Define http server
	s := &http.Server{
		Addr:    ":" + port,
		Handler: m,
	}

	go func() {
		log.Println("🚀 Server is running on port", s.Addr)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}

	}()

	sigChan, closeChan := createChannel()
	defer closeChan()

	log.Println("Shutting down server...", <-sigChan)

	tc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.Shutdown(tc)
}

func createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}
