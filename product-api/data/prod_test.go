package data

import "testing"

func TestCheck(t *testing.T) {
	p := &Product{
		Name:  "kumar",
		Price: 23,
		SKU:   "abs-shd-cdd",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
