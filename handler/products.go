package handler

import (
	"log"
	"net/http"
	"regexp"

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
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// PUT -> request
	if r.Method == http.MethodPut {
		p.l.Println("PUT :", r.URL.Path)

		//expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
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

// POST ->
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Printf("Error decoding JSON: %v", err)
		http.Error(rw, "Unable to unmarshal json data to object", http.StatusBadRequest)
		return
	}

	// p.l.Printf("Product: %#v", prod)
	data.AddProduct(prod)
}
