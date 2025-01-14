package handler

import (
	"log"
	"net/http"

	"github.com/abhiraj-ku/micro_serGO/data"
)

type Products struct {
	l *log.Logger
}

// contructor to init the Product struct
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	//catch if no method satisfies the GET
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

// getProducts returns the products from the data store

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")
	productList := data.GetProducts()

	err := productList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal your data into json", http.StatusInternalServerError)
	}

}
