package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newCategory)
}

// PUT localhost:8080/api/categories/{id}
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := r.PathValue("id")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	category.ID = id

	err = h.service.Update(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/categories/123 -> ID = 123
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := r.PathValue("id")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted successfully",
	})
}
