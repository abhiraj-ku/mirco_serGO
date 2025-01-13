package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// reqeusts to the path /goodbye with be handled by this function
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// any other request will be handled by this function
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello Handler")

		// read the body
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// write the response
		fmt.Fprintf(rw, "Hello %s", b)
	})

	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
