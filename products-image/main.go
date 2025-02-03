package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abhiraj-ku/micro_serGO/files"
	"github.com/abhiraj-ku/micro_serGO/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("This is products image service")

	mux := mux.NewRouter()

	// file handlers
	stor, err := files.NewLocal(*basePath, 1024*1000*5)
	if err != nil {
		log.Fatal("Unable to create storage", "error", err)
		os.Exit(1)
	}
	// Create the handler for storage
	fh := handlers.NewFile(stor, log.Logger)

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	// problem with FileServer is that it is dumb
	ph := mux.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	// get files
	gh := mux.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(*basePath))),
	)

	// Create a server
	server := http.Server{
		Addr:         ":4040",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	// Start the server as goroutines

	go func() {
		log.Println("Starting the server on port 4040")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block untill a request is recieved
	sig := <-c
	log.Print("shutting server down", "signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
