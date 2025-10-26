package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

var products []models.Product

func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ID = uuid.New().String()
	products = append(products, p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
