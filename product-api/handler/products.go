// Package classification of Product API
//
// DOCUMENTATION USING SWAGGER
//
// Schemas: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
//	 - application/json
//
// Produces:
//  - application/json
// swagger:meta

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
	err = data.UpdateProduct(id)
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
			p.l.Printf("[Error] decoding JSON: %v", err)
			http.Error(rw, "Unable to unmarshal json data to object", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Printf("[Error] Validating the input: %v", err)
			http.Error(rw, "Unable to unmarshal json data to object", http.StatusBadRequest)
			return

		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
