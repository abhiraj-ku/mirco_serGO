package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"createdOn"`
	UpdatedOn   string  `json:"-"`
}

// Validator function
func (p *Product) Validate() error {
	validate := validator.New()

	// custom function (here sku check)
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

// custom validator func for SKU
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+-`)
	validSKUs := re.FindAllString(fl.Field().String(), -1)

	//TODO: Fix the "sku" tag test custom validator message failed
	if len(validSKUs) != 1 {
		return false
	}

	return true

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
func UpdateProduct(p Product) error {
	i := findProdById(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = &p

	return nil
}

func DeleteProduct(id int) error {
	i := findProdById(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

// find index by prod id

func findProdById(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}
