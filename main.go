package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/d3z41k/rest-boilerplate/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpPort := os.Getenv("HTTP_PORT")

	handler := server.NewRouter()
	srv := &http.Server{
		Addr:         ":" + httpPort,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server
	go srv.ListenAndServe()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
