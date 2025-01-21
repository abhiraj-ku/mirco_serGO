package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/abhiraj-ku/micro_serGO/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

// contructor to init the Product struct
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}
// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	// PUT -> request
// 	if r.Method == http.MethodPut {
// 		p.l.Println("PUT :", r.URL.Path)

// 		// //expect the id in the URI
// 		// reg := regexp.MustCompile(`/([0-9]+)`)
// 		// g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		// if len(g) != 1 {
// 		// 	p.l.Println("Invalid URI more than one id")
// 		// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 		// 	return
// 		// }

// 		// if len(g[0]) != 2 {
// 		// 	p.l.Println("Invalid URI more than one id")
// 		// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 		// 	return
// 		// }

// 		// idString, err := strconv.ParseInt(g[0][1], 10, 0)
// 		// if err != nil {
// 		// 	http.Error(rw, "Error converting idString to int", http.StatusBadRequest)
// 		// 	return
// 		// }

// 		// fmt.Println("ID :", idString)\

// 		// Simpler method to parse ID params from url Path

// 		urlParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
// 		if len(urlParts) != 1 {
// 			p.l.Println("Cannot parse URL path ,Invalid URl")
// 			http.Error(rw, "Invalid URL", http.StatusBadRequest)
// 			return
// 		}
// 		id, err := strconv.Atoi(urlParts[0])
// 		if err != nil {
// 			p.l.Println("Cannot parse URL path ,Invalid URl")
// 			http.Error(rw, "Invalid ID", http.StatusBadRequest)
// 			return
// 		}

// 		p.updateProducts(id, rw, r)
// 		return

// 	}

// 	//catch if no method satisfies the GET
// 	rw.WriteHeader(http.StatusMethodNotAllowed)

// }

// getProducts returns the products from the data store

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")
	productList := data.GetProducts()

	err := productList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal your data into json", http.StatusInternalServerError)
	}

}

// POST ->addProduct
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Printf("Error decoding JSON: %v", err)
		http.Error(rw, "Unable to unmarshal json data to object", http.StatusBadRequest)
		return
	}

	// p.l.Printf("Product: %#v", prod)
	data.AddProduct(&prod)
}

// PUT ->updateProducts

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = prod.FromJSON(r.Body)
	if err != nil {
		p.l.Printf("Error decoding JSON: %v", err)
		http.Error(rw, "Unable to unmarshal json data to object", http.StatusBadRequest)
		return
	}
	// Update the productList
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found for update", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found for update", http.StatusInternalServerError)
		return

	}
}

// Middleware to parse json and vice versa
type KeyProduct struct{}

func (p *Products) ValidateInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Printf("Error decoding JSON: %v", err)
			http.Error(rw, "Unable to unmarshal json data to object", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
