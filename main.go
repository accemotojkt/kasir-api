package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/handlers"
	"net/http"
)

func main() {
	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	http.HandleFunc("GET /api/products", handlers.ListProducts)
	http.HandleFunc("POST /api/products", handlers.CreateProduct)
	http.HandleFunc("GET /api/products/{id}", handlers.GetProductByID)
	http.HandleFunc("PUT /api/products/{id}", handlers.UpdateProduct)
	http.HandleFunc("DELETE /api/products/{id}", handlers.DeleteProduct)

	http.HandleFunc("GET /api/categories", handlers.ListCategories)
	http.HandleFunc("POST /api/categories", handlers.CreateCategory)
	http.HandleFunc("GET /api/categories/{id}", handlers.GetCategoryByID)
	http.HandleFunc("PUT /api/categories/{id}", handlers.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", handlers.DeleteCategory)

	fmt.Println("Server Running di Localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal Running Server")
	}

}
