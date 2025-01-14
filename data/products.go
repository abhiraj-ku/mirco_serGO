package data

import (
	"encoding/json"
	"io"
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

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}
