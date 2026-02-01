package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"net/http"
	"strconv"
)

// In-memory storage (sementara, nanti ganti database)
var products = []models.Product{
	{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
	{ID: 3, Name: "kecap", Price: 12000, Stock: 20},
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// masukkin data ke dalam variable products
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newProduct)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/products/123 -> ID = 123
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Cari produk dengan ID tersebut
	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Product not found", http.StatusNotFound)
}

// PUT localhost:8080/api/products/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := r.PathValue("id")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop products, cari id, ganti sesuai data dari request
	for i := range products {
		if products[i].ID == id {
			updateProduct.ID = id
			products[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := r.PathValue("id")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// loop products cari ID, dapet index yang mau dihapus
	for i, p := range products {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			products = append(products[:i], products[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
