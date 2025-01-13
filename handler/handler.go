package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Note: ResponseWriter is an Interface and Request is a struct
// so we use ResponseWriter without a pointer and Request as a pointer
// to the http.Request struct

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handling hello request")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Printf("Error reading body: %v", err)
		http.Error(rw, "Cannot read body contents", http.StatusBadRequest)
		return
	}

	h.l.Printf("Responding with name: %s", d)
	fmt.Fprintf(rw, "Hello %s ", d)
}
