package main

import (
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	// Добавляем обработку статических файлов
	fsUploads := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fsUploads))

	fsTemplates := http.FileServer(http.Dir("./templates"))
	mux.Handle("/templates/", http.StripPrefix("/templates/", fsTemplates))

	handler := corsMiddleware(mux)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
