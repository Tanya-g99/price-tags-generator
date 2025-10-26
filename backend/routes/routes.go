package routes

import (
	"backend/handlers"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	// Продукты
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListProducts(w, r)
		case http.MethodPost:
			handlers.AddProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Ценники
	mux.HandleFunc("/api/pricetags", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListPriceTags(w, r)
		case http.MethodPost:
			handlers.AddPriceTag(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Шаблоны
	mux.HandleFunc("/api/templates-pricetags", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListPriceTagTemplates(w, r)
		case http.MethodPost:
			handlers.AddPriceTagTemplate(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
