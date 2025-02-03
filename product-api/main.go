package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abhiraj-ku/micro_serGO/handler"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "products-api: ", log.LstdFlags)
	hp := handler.NewProducts(l)
	mux := mux.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", hp.GetProducts)

	// POST ->
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/add", hp.AddProduct)

	putRouter := mux.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", hp.UpdateProducts)
	putRouter.Use(hp.ValidateInput)

	// CORS header
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://hellobunny.com:3000"}))

	// mux.Handle("/", hp)

	// err := http.ListenAndServe(":9090", mux)
	// log.Fatal(err)

	// custom server in golang
	log.Println("Server is up and running...")
	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(mux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// err := s.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	go func() {
		l.Println("starting the server on port 9090...")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}

	}()
	// Graceful shutdown

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	// block until a signal is recieved

	sig := <-c
	log.Println("Got signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}
