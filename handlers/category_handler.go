package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"net/http"
	"strconv"
)

// In-memory storage (sementara, nanti ganti database)
var categories = []models.Category{
	{ID: 1, Name: "Mie Instant", Description: "Mie"},
	{ID: 2, Name: "Minuman", Description: "Minum Botol"},
}

func ListCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// masukkin data ke dalam variable categories
	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newCategory)
}

// PUT localhost:8080/api/categories/{id}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := r.PathValue("id")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop categories, cari id, ganti sesuai data dari request
	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id
			categories[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategory)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/categories/123 -> ID = 123
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// Cari categories dengan ID tersebut
	for _, p := range categories {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Category not found", http.StatusNotFound)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := r.PathValue("id")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// loop categories cari ID, dapet index yang mau dihapus
	for i, p := range categories {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
