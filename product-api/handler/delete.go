package handler

import (
	"net/http"
	"strconv"

	"github.com/abhiraj-ku/micro_serGO/data"
	"github.com/gorilla/mux"
)

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle delete PRoduct", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Can't delete this as prod not found", http.StatusInternalServerError)
		return
	}
}
