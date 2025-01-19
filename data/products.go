package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"createdOn"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Herbal Tea",
		Description: "Natural tea from the gardens of Assam",
		Price:       9000,
		SKU:         "bcmc890",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	}, &Product{
		ID:          2,
		Name:        "Herbal Coffee",
		Description: "Natural tea from the gardens of Assam",
		Price:       9000,
		SKU:         "bcmc890",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

type Products []*Product

// Encode the objects to json (marshal the data)
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Decode the json from body to go's Object (struct) (unmarshal the data)
func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	if err := dec.Decode(p); err != nil {
		log.Fatal(err)

	}
	return nil
}

func GetProducts() Products {
	return productList
}

func getNextID() int {
	lp := productList[len(productList)-1]

	return lp.ID + 1
}
func AddProduct(p *Product) {
	p.ID = getNextID()

	productList = append(productList, p)

}

// Error format var for findProd error
var ErrProductNotFound = fmt.Errorf("Product Not found")

// find product method to find product based on the given id
func findProd(id int) (*Product, int, error) {
	for i, product := range productList {
		if product.ID == id {
			return product, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func UpdateProduct(id int, p *Product) error {
	_, indx, err := findProd(id)
	if err != nil {
		return fmt.Errorf("product update failed: %w", err)
	}
	p.ID = id
	productList[indx] = p
	return nil

}
