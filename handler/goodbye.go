package handler

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

// This is temp handler for demo purpose

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Bye from bye handler"))
}
